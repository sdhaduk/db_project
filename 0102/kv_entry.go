package db0102

import (
	"encoding/binary"
	"io"
)

type Entry struct {
	key []byte
	val []byte
}

func (ent *Entry) Encode() []byte {
	res := make([]byte, 0)

	key_buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(key_buf, uint32(len(ent.key)))

	val_buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(val_buf, uint32(len(ent.val)))

	res = append(res, key_buf...)
	res = append(res, val_buf...)
	res = append(res, ent.key...)
	res = append(res, ent.val...)
	
	return res
}

func (ent *Entry) Decode(r io.Reader) error {
	size_buf := make([]byte, 4)
	_, err := r.Read(size_buf)
	if err != nil {
		return err
	}

	key_size := binary.LittleEndian.Uint32(size_buf)

	_, err = r.Read(size_buf)
	if err != nil {
		return err
	}

	val_size := binary.LittleEndian.Uint32(size_buf)

	key_buf := make([]byte, key_size)
	val_buf := make([]byte, val_size)

	_, err = r.Read(key_buf)
	if err != nil {
		return err
	}

	_, err = r.Read(val_buf)
	if err != nil {
		return err
	}

	ent.key = key_buf
	ent.val = val_buf
	return nil
}
