// Code generated by fastssz. DO NOT EDIT.
package p2p

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the MsgAddr object
func (m *MsgAddr) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the MsgAddr object to a target array
func (m *MsgAddr) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Addr'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(m.Addr) * 64

	// Field (0) 'Addr'
	if len(m.Addr) > 32 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(m.Addr); ii++ {
		dst = append(dst, m.Addr[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the MsgAddr object
func (m *MsgAddr) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Addr'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Field (0) 'Addr'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 64, 32)
		if err != nil {
			return err
		}
		m.Addr = make([][64]byte, num)
		for ii := 0; ii < num; ii++ {
			copy(m.Addr[ii][:], buf[ii*64:(ii+1)*64])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MsgAddr object
func (m *MsgAddr) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Addr'
	size += len(m.Addr) * 64

	return
}

// HashTreeRoot ssz hashes the MsgAddr object
func (m *MsgAddr) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the MsgAddr object with a hasher
func (m *MsgAddr) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Addr'
	{
		if len(m.Addr) > 32 {
			err = ssz.ErrListTooBig
			return
		}
		subIndx := hh.Index()
		for _, i := range m.Addr {
			hh.Append(i[:])
		}
		numItems := uint64(len(m.Addr))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(32, numItems, 32))
	}

	hh.Merkleize(indx)
	return
}