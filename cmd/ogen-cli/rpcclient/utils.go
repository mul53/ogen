package rpcclient

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	bls_interface "github.com/olympus-protocol/ogen/pkg/bls/interface"
	"strconv"
	"time"

	"github.com/olympus-protocol/ogen/api/proto"
	"github.com/olympus-protocol/ogen/pkg/bls"
)

func (c *RPCClient) stopProposer(args []string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.utils.StopProposer(ctx, &proto.Empty{})
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *RPCClient) startProposer([]string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	res, err := c.utils.StartProposer(ctx, &proto.Empty{})
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *RPCClient) submitRawData(args []string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if len(args) < 2 {
		return "", errors.New("Usage: submitrawdata <raw_data> <type>")
	}
	req := &proto.RawData{Data: args[0], Type: args[1]}
	res, err := c.utils.SubmitRawData(ctx, req)
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *RPCClient) genKeyPair(raw bool) (string, error) {
	blsKeyPair := bls.CurrImplementation.RandKey()
	var res bls_interface.KeyPair
	if raw {
		res = bls_interface.KeyPair{
			Public:  hex.EncodeToString(blsKeyPair.PublicKey().Marshal()),
			Private: hex.EncodeToString(blsKeyPair.Marshal()),
		}
	} else {
		addr, err := blsKeyPair.PublicKey().ToAccount()
		if err != nil {
			return "", errors.New("unable to encode public key to bech32")
		}
		wif, err := blsKeyPair.ToWIF()
		if err != nil {
			return "", errors.New("unable to encode private key to bech32")
		}
		res = bls_interface.KeyPair{
			Public:  addr,
			Private: wif,
		}
	}

	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *RPCClient) genValidatorKey(args []string) (out string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	amount := 0
	if len(args) < 1 {
		return "", errors.New("Usage: genvalidatorkey <keys>")
	}
	amount, err = strconv.Atoi(args[0])
	if err != nil {
		return out, err
	}
	req := &proto.GenValidatorKeys{
		Keys: uint64(amount),
	}
	res, err := c.utils.GenValidatorKey(ctx, req)
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *RPCClient) decodeRawTransaction(args []string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if len(args) < 1 {
		return "", errors.New("Usage: decoderawtransaction <raw_transaction>")
	}
	req := &proto.RawData{Data: args[0]}
	res, err := c.utils.DecodeRawTransaction(ctx, req)
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *RPCClient) decodeRawBlock(args []string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if len(args) < 1 {
		return "", errors.New("Usage: decoderawblock <raw_block>")
	}
	req := &proto.RawData{Data: args[0]}
	res, err := c.utils.DecodeRawBlock(ctx, req)
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
