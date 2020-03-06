package blockdb

import (
	"bytes"
	"github.com/olympus-protocol/ogen/p2p"
	"github.com/olympus-protocol/ogen/primitives"
	"github.com/olympus-protocol/ogen/utils/chainhash"
	"os"
	"testing"
	"time"
)

var (
	testBlockLocator = BlockLocation{
		FileNum:     100,
		BlockOffset: 100000,
		BlockSize:   10000000,
	}
	merkleRootTest = chainhash.Hash([chainhash.HashSize]byte{
		0xfc, 0x4b, 0x8c, 0xb9, 0x03, 0xae, 0xd5, 0x4e,
		0x11, 0xe1, 0xae, 0x8a, 0x5b, 0x7a, 0xd0, 0x97,
		0xad, 0xe3, 0x49, 0x88, 0xa8, 0x45, 0x00, 0xad,
		0x2d, 0x80, 0xe4, 0xd1, 0xf5, 0xbc, 0xc9, 0x5d,
	})
	TxTest = primitives.Tx{
		TxVersion: 1,
		TxType:    primitives.Coins,
		TxAction:  primitives.Transfer,
		Time:      time.Unix(1572830409, 0).Unix(),
	}
	TestBlock = primitives.Block{
		Header: primitives.BlockHeader{
			Version:       1,
			PrevBlockHash: chainhash.Hash{},
			MerkleRoot:    merkleRootTest,
			Timestamp:     time.Unix(0x5A3BB72B, 0),
		},
		Txs:       []primitives.Tx{TxTest, TxTest, TxTest, TxTest, TxTest, TxTest, TxTest, TxTest, TxTest, TxTest},
		Signature: [96]byte{},
	}
)

func Test_BlockLocatorSerialize(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	err := testBlockLocator.Serialize(buf)
	if err != nil {
		t.Errorf("block locator serialize, unable to serialize")
		return
	}
	var newBlockLoc BlockLocation
	err = newBlockLoc.Deserialize(buf)
	if err != nil {
		t.Errorf("block locator serialize, unable to deserialize")
		return
	}
	if newBlockLoc.FileNum != testBlockLocator.FileNum {
		t.Errorf("block locator serialize file num doesn't match")
		return
	}
	if newBlockLoc.BlockOffset != testBlockLocator.BlockOffset {
		t.Errorf("block locator serialize block offset doesn't match")
		return
	}
	if newBlockLoc.BlockSize != testBlockLocator.BlockSize {
		t.Errorf("block locator serialize block size doesn't match")
		return
	}
}

func TestRawBlockDB_Write(t *testing.T) {
	for i := 0; i < 1000; i++ {
		buf := bytes.NewBuffer([]byte{})
		err := TestBlock.Encode(buf)
		if err != nil {
			t.Errorf("block write error: %v", err.Error())
			return
		}
		loc, err := blockDB.rawBlockDb.write(buf.Bytes())
		if err != nil {
			t.Errorf("block write error: %v", err.Error())
			return
		}
		hash := TestBlock.Header.Hash()
		blockBytes, err := blockDB.rawBlockDb.read(hash, loc)
		if err != nil {
			t.Errorf("block write read: %v", err.Error())
			return
		}
		newBlockBuf := bytes.NewBuffer(blockBytes)
		var block p2p.MsgBlock
		err = block.Decode(newBlockBuf)
		if err != nil {
			t.Errorf("block write read: %v", err.Error())
			return
		}
		testHash := TestBlock.Header.Hash()
		newBlockHash := block.Header.Hash()
		if testHash != newBlockHash {
			t.Errorf("block write read: hashes doesn't match")
			return
		}
	}
	finish()
}

func finish() {
	_ = os.RemoveAll("./blocks")
	_ = os.RemoveAll("./db")
}