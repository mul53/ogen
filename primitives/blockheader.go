package primitives

import (
	"bytes"
	"github.com/olympus-protocol/ogen/utils/chainhash"
	"github.com/olympus-protocol/ogen/utils/serializer"
	"io"
	"time"
)

var MaxBlockHeaderBytes = 76

type BlockHeader struct {
	Version       int32
	Nonce         int32
	MerkleRoot    chainhash.Hash
	PrevBlockHash chainhash.Hash
	Timestamp     time.Time
}

func (bh *BlockHeader) Serialize(w io.Writer) error {
	sec := uint32(bh.Timestamp.Unix())
	err := serializer.WriteElements(w, bh.Version, bh.Nonce, bh.MerkleRoot, bh.PrevBlockHash, sec)
	if err != nil {
		return err
	}
	return nil
}

func (bh *BlockHeader) Deserialize(r io.Reader) error {
	err := serializer.ReadElements(r, &bh.Version, &bh.Nonce, &bh.MerkleRoot, &bh.PrevBlockHash, (*serializer.Uint32Time)(&bh.Timestamp))
	if err != nil {
		return err
	}
	return nil
}

func (bh *BlockHeader) Hash() chainhash.Hash {
	buf := bytes.NewBuffer([]byte{})
	_ = bh.Serialize(buf)
	return chainhash.DoubleHashH(buf.Bytes())
}

func NewBlockHeader(version int32, prevBlock chainhash.Hash, nonce int32, merkle chainhash.Hash, time time.Time) *BlockHeader {
	return &BlockHeader{
		Version:       version,
		Nonce:         nonce,
		MerkleRoot:    merkle,
		PrevBlockHash: prevBlock,
		Timestamp:     time,
	}
}
