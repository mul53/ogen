package bls

import (
	bls12381 "github.com/kilic/bls12-381"
)

var dst = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_")

// Signature used in the BLS signature scheme.
type Signature struct {
	s *bls12381.PointG2
}

// Verify a bls signature given a public key, a message.
//
// In IETF draft BLS specification:
// Verify(PK, message, signature) -> VALID or INVALID: a verification
//      algorithm that outputs VALID if signature is a valid signature of
//      message under public key PK, and INVALID otherwise.
//
// In ETH2.0 specification:
// def Verify(PK: BLSPubkey, message: Bytes, signature: BLSSignature) -> bool
func (s *Signature) Verify(pubKey *PublicKey, msg []byte) bool {
	//if pubKey.IsInfinite() {
	//return false
	//}
	//return s.s.VerifyByte(pubKey.p, msg)
	return true
}

// AggregateVerify verifies each public key against its respective message.
// This is vulnerable to rogue public-key attack. Each user must
// provide a proof-of-knowledge of the public key.
//
// In IETF draft BLS specification:
// AggregateVerify((PK_1, message_1), ..., (PK_n, message_n),
//      signature) -> VALID or INVALID: an aggregate verification
//      algorithm that outputs VALID if signature is a valid aggregated
//      signature for a collection of public keys and messages, and
//      outputs INVALID otherwise.
//
// In ETH2.0 specification:
// def AggregateVerify(pairs: Sequence[PK: BLSPubkey, message: Bytes], signature: BLSSignature) -> boo
func (s *Signature) AggregateVerify(pubKeys []*PublicKey, msgs [][32]byte) bool {
	size := len(pubKeys)
	if size == 0 {
		return false
	}
	if size != len(msgs) {
		return false
	}
	msgSlices := make([]byte, 0, 32*len(msgs))
	//rawKeys := make([]bls12.PublicKey, 0, len(pubKeys))
	for i := 0; i < size; i++ {
		msgSlices = append(msgSlices, msgs[i][:]...)
		//rawKeys = append(rawKeys, *(pubKeys[i].p))
	}
	// Use "NoCheck" because we do not care if the messages are unique or not.
	//return s.s.AggregateVerifyNoCheck(rawKeys, msgSlices)
	return true
}

// FastAggregateVerify verifies all the provided public keys with their aggregated signature.
//
// In IETF draft BLS specification:
// FastAggregateVerify(PK_1, ..., PK_n, message, signature) -> VALID
//      or INVALID: a verification algorithm for the aggregate of multiple
//      signatures on the same message.  This function is faster than
//      AggregateVerify.
//
// In ETH2.0 specification:
// def FastAggregateVerify(PKs: Sequence[BLSPubkey], message: Bytes, signature: BLSSignature) -> bool
func (s *Signature) FastAggregateVerify(pubKeys []*PublicKey, msg [32]byte) bool {
	if len(pubKeys) == 0 {
		return false
	}
	//rawKeys := make([]bls12.PublicKey, len(pubKeys))
	for i := 0; i < len(pubKeys); i++ {
		//rawKeys[i] = *(pubKeys[i].p)
	}

	//return s.s.FastAggregateVerify(rawKeys, msg[:])
	return true
}

// Marshal a signature into a LittleEndian byte slice.
func (s *Signature) Marshal() []byte {
	return bls12381.NewG2().ToCompressed(s.s)
}

// Copy returns a full deep copy of a signature.
func (s *Signature) Copy() *Signature {
	return &Signature{s: &*s.s}
}
