package db0105

import (
	"encoding/binary"
	"errors"
	"hash/crc32"
	"io"
)

type Entry struct {
	key     []byte
	val     []byte
	deleted bool
}

var ErrBadSum = errors.New("bad checksum")

func (ent *Entry) Encode() []byte {
	data := make([]byte, 4+4+4+1+len(ent.key)+len(ent.val))
	binary.LittleEndian.PutUint32(data[4:8], uint32(len(ent.key)))
	copy(data[13:], ent.key)
	if ent.deleted {
		data[12] = 1
	} else {
		binary.LittleEndian.PutUint32(data[8:12], uint32(len(ent.val)))
		copy(data[13+len(ent.key):], ent.val)
	}
	binary.LittleEndian.PutUint32(data[0:4], crc32.ChecksumIEEE(data[4:]))

	return data
}

func (ent *Entry) Decode(r io.Reader) error {
	var checksum_bytes [4]byte
	if _, err := io.ReadFull(r, checksum_bytes[:]); err != nil {
		return err
	}
	checksum := binary.LittleEndian.Uint32(checksum_bytes[:])

	var header [9]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return err
	}
	klen := int(binary.LittleEndian.Uint32(header[0:4]))
	vlen := int(binary.LittleEndian.Uint32(header[4:8]))
	deleted := header[8]

	data := make([]byte, klen+vlen)
	if _, err := io.ReadFull(r, data); err != nil {
		return err
	}

	combined := append(header[:], data...)
	checksum_calc := crc32.ChecksumIEEE(combined)

	if checksum != checksum_calc {
		return ErrBadSum
	}

	ent.key = data[:klen]
	if deleted != 0 {
		ent.deleted = true
	} else {
		ent.deleted = false
		ent.val = data[klen:]
	}

	return nil
}
// QzBQWVJJOUhU https://trialofcode.org/
