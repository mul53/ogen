package wallet

import (
	"bytes"
	"fmt"

	"github.com/olympus-protocol/ogen/bls"
	"github.com/olympus-protocol/ogen/primitives"
	"github.com/olympus-protocol/ogen/utils/bech32"
	"github.com/olympus-protocol/ogen/utils/chainhash"
)

// GetBalance gets the balance of an address or of the wallet address.
func (b *Wallet) GetBalance(addr string) (uint64, error) {
	if addr == "" {
		addr = b.info.address
	}
	_, pkh, err := bech32.Decode(addr)
	if err != nil {
		return 0, err
	}
	if len(pkh) != 20 {
		return 0, fmt.Errorf("expecting address to be 20 bytes, but got %d bytes", len(pkh))
	}
	var pkhBytes [20]byte
	copy(pkhBytes[:], pkh)
	out, ok := b.chain.State().TipState().UtxoState.Balances[pkhBytes]
	if !ok {
		return 0, nil
	}
	return out, nil
}

// GetAddressRaw returns the pubkey hash of the wallet.
func (b *Wallet) GetAddressRaw() ([20]byte, error) {
	_, pkh, err := bech32.Decode(b.info.address)
	if err != nil {
		return [20]byte{}, err
	}
	if len(pkh) != 20 {
		return [20]byte{}, fmt.Errorf("expecting address to be 20 bytes, but got %d bytes", len(pkh))
	}
	var pkhBytes [20]byte
	copy(pkhBytes[:], pkh)
	return pkhBytes, nil
}

func (b *Wallet) broadcastTx(payload *primitives.CoinPayload) {
	buf := bytes.NewBuffer([]byte{})
	err := payload.Encode(buf)
	if err != nil {
		b.log.Errorf("error encoding transaction: %s", err)
		return
	}
	if err := b.txTopic.Publish(b.ctx, buf.Bytes()); err != nil {
		b.log.Errorf("error broadcasting transaction: %s", err)
	}
}

func (b *Wallet) broadcastDeposit(deposit *primitives.Deposit) {
	buf := bytes.NewBuffer([]byte{})
	err := deposit.Encode(buf)
	if err != nil {
		b.log.Errorf("error encoding transaction: %s", err)
		return
	}
	if err := b.depositTopic.Publish(b.ctx, buf.Bytes()); err != nil {
		b.log.Errorf("error broadcasting transaction: %s", err)
	}
}

// StartValidator signs a validator deposit.
func (b *Wallet) StartValidator(authentication []byte, validatorPrivBytes [32]byte) (*primitives.Deposit, error) {
	priv, err := b.unlockIfNeeded(authentication)
	if err != nil {
		return nil, err
	}
	pub := priv.PublicKey()

	validatorPriv, err := bls.SecretKeyFromBytes(validatorPrivBytes[:])
	if err != nil {
		return nil, err
	}

	validatorPub := validatorPriv.PublicKey()
	validatorPubBytes := validatorPub.Marshal()
	validatorPubHash := chainhash.HashH(validatorPubBytes[:])

	validatorProofOfPossession := validatorPriv.Sign(validatorPubHash[:])

	addr, err := b.GetAddressRaw()
	if err != nil {
		return nil, err
	}

	depositData := &primitives.DepositData{
		PublicKey:         *validatorPub,
		ProofOfPossession: *validatorProofOfPossession,
		WithdrawalAddress: addr,
	}

	buf := bytes.NewBuffer([]byte{})

	if err := depositData.Encode(buf); err != nil {
		return nil, err
	}

	depositHash := chainhash.HashH(buf.Bytes())

	depositSig := priv.Sign(depositHash[:])

	deposit := &primitives.Deposit{
		PublicKey: *pub,
		Signature: *depositSig,
		Data:      *depositData,
	}

	currentState := b.chain.State().TipState()

	if err := b.actionMempool.AddDeposit(deposit, currentState); err != nil {
		return nil, err
	}

	b.broadcastDeposit(deposit)

	return deposit, nil
}

func (b *Wallet) broadcastExit(exit *primitives.Exit) {
	buf := bytes.NewBuffer([]byte{})
	err := exit.Encode(buf)
	if err != nil {
		b.log.Errorf("error encoding transaction: %s", err)
		return
	}
	if err := b.exitTopic.Publish(b.ctx, buf.Bytes()); err != nil {
		b.log.Errorf("error broadcasting transaction: %s", err)
	}
}

// ExitValidator submits an exit transaction for a certain validator.
func (w *Wallet) ExitValidator(authentication []byte, validatorPubKey [48]byte) (*primitives.Exit, error) {
	priv, err := w.unlockIfNeeded(authentication)
	if err != nil {
		return nil, err
	}

	validatorPub, err := bls.PublicKeyFromBytes(validatorPubKey[:])
	if err != nil {
		return nil, err
	}

	currentState := w.chain.State().TipState()

	pub := priv.PublicKey()

	msg := fmt.Sprintf("exit %x", validatorPub.Marshal())
	msgHash := chainhash.HashH([]byte(msg))

	sig := priv.Sign(msgHash[:])

	exit := &primitives.Exit{
		ValidatorPubkey: *validatorPub,
		WithdrawPubkey:  *pub,
		Signature:       *sig,
	}

	if err := w.actionMempool.AddExit(exit, currentState); err != nil {
		return nil, err
	}

	w.broadcastExit(exit)

	return exit, nil
}

// SendToAddress sends an amount to an address with the given password and parameters.
func (b *Wallet) SendToAddress(authentication []byte, to string, amount uint64) (*chainhash.Hash, error) {
	priv, err := b.unlockIfNeeded(authentication)
	if err != nil {
		return nil, err
	}

	_, data, err := bech32.Decode(to)
	if err != nil {
		return nil, err
	}

	if len(data) != 20 {
		return nil, fmt.Errorf("invalid address")
	}

	var toPkh [20]byte

	copy(toPkh[:], data)

	pub := priv.PublicKey()

	b.lastNonceLock.Lock()
	b.info.lastNonce++
	nonce := b.info.lastNonce
	b.lastNonceLock.Unlock()

	payload := &primitives.CoinPayload{
		To:            toPkh,
		FromPublicKey: *pub,
		Amount:        amount,
		Nonce:         nonce,
		Fee:           1,
	}

	sigMsg := payload.SignatureMessage()
	sig := priv.Sign(sigMsg[:])

	payload.Signature = *sig

	tx := &primitives.Tx{
		TxType:    primitives.TxCoins,
		TxVersion: 0,
		Payload:   payload,
	}

	currentState := b.chain.State().TipState()

	if err := b.mempool.Add(*payload, &currentState.UtxoState); err != nil {
		return nil, err
	}

	b.broadcastTx(payload)

	txHash := tx.Hash()

	return &txHash, nil
}