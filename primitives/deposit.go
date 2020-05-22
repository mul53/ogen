package primitives

import (
	"bytes"
	"io"

	"github.com/olympus-protocol/ogen/bls"
	"github.com/olympus-protocol/ogen/utils/chainhash"
)

// Deposit is a deposit a user can submit to queue as a validator.
type Deposit struct {
	// PublicKey is the public key of the address that is depositing.
	PublicKey bls.PublicKey

	// Signature is the signature signing the deposit data.
	Signature bls.Signature

	// Data is the data that describes the new validator.
	Data DepositData
}

// Encode encodes the deposit to the given writer.
func (d *Deposit) Encode(w io.Writer) error {
	sigBytes := d.Signature.Marshal()
	pubBytes := d.PublicKey.Marshal()

	if _, err := w.Write(pubBytes[:]); err != nil {
		return err
	}
	if _, err := w.Write(sigBytes[:]); err != nil {
		return err
	}
	if err := d.Data.Encode(w); err != nil {
		return err
	}
	return d.Data.Encode(w)
}

// Decode decodes the deposit to the given writer.
func (d *Deposit) Decode(r io.Reader) error {
	sigBytes := make([]byte, 96)
	pubBytes := make([]byte, 48)
	if _, err := r.Read(pubBytes[:]); err != nil {
		return err
	}
	if _, err := r.Read(sigBytes[:]); err != nil {
		return err
	}
	if err := d.Data.Decode(r); err != nil {
		return err
	}
	sig, err := bls.SignatureFromBytes(sigBytes)
	if err != nil {
		return err
	}
	pub, err := bls.PublicKeyFromBytes(pubBytes)
	if err != nil {
		return err
	}
	d.Signature = *sig
	d.PublicKey = *pub
	return d.Data.Decode(r)
}

// Hash calculates the hash of the deposit
func (d *Deposit) Hash() chainhash.Hash {
	buf := bytes.NewBuffer([]byte{})
	_ = d.Encode(buf)
	return chainhash.HashH(buf.Bytes())
}

// DepositData is the part of the deposit that is signed
type DepositData struct {
	// PublicKey is the key used for the validator.
	PublicKey bls.PublicKey

	// ProofOfPossession is the public key signed by the private key to prove that you
	// own the address and prevent rogue public-key attacks.
	ProofOfPossession bls.Signature

	// WithdrawalAddress is the address to withdraw to.
	WithdrawalAddress [20]byte
}

// Encode encodes the deposit data to a writer.
func (dd *DepositData) Encode(w io.Writer) error {
	pubBytes := dd.PublicKey.Marshal()
	proofOfPossessionBytes := dd.ProofOfPossession.Marshal()

	if _, err := w.Write(pubBytes[:]); err != nil {
		return err
	}
	if _, err := w.Write(proofOfPossessionBytes[:]); err != nil {
		return err
	}
	if _, err := w.Write(dd.WithdrawalAddress[:]); err != nil {
		return err
	}

	return nil
}

// Decode decodes the deposit data from a reader.
func (dd *DepositData) Decode(r io.Reader) error {
	pubBytes := make([]byte, 48)
	sigBytes := make([]byte, 48)
	var withdrawalAddress [20]byte

	if _, err := io.ReadFull(r, pubBytes[:]); err != nil {
		return err
	}
	if _, err := io.ReadFull(r, sigBytes[:]); err != nil {
		return err
	}
	if _, err := io.ReadFull(r, withdrawalAddress[:]); err != nil {
		return err
	}

	sig, err := bls.SignatureFromBytes(sigBytes)
	if err != nil {
		return err
	}
	pub, err := bls.PublicKeyFromBytes(pubBytes)
	if err != nil {
		return err
	}
	dd.WithdrawalAddress = withdrawalAddress
	dd.PublicKey = *pub
	dd.ProofOfPossession = *sig

	return nil
}