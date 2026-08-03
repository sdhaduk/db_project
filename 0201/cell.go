package db0201

import "encoding/binary"

type CellType uint8

const (
	TypeI64 CellType = 1
	TypeStr CellType = 2
)

type Cell struct {
	Type CellType
	I64  int64
	Str  []byte
}

func (cell *Cell) Encode(toAppend []byte) []byte {
	switch cell.Type {
	case TypeI64:
		encoding := binary.LittleEndian.AppendUint64(toAppend, uint64(cell.I64))
		return encoding
	
	case TypeStr:
		encoding := make([]byte, 4 + len(cell.Str))
		binary.LittleEndian.PutUint32(encoding, uint32(len(cell.Str)))
		copy(encoding[4:], cell.Str)
		encoding = append(toAppend, encoding...)
		return encoding
	
	default:
		panic("unkown cell type")
	}
}

func (cell *Cell) Decode(data []byte) (rest []byte, err error) {
	switch cell.Type {
	case TypeI64:
		cell.I64 = int64(binary.LittleEndian.Uint64(data[:8]))
		return data[8:], nil
	
	case TypeStr:
		len_data := binary.LittleEndian.Uint32(data[:4])
		str_data := data[4:4+len_data]
		cell.Str = str_data
		return data[4+len_data:], nil
	
	default:
		panic("unknown cell type")

	}
}

// QzBQWVJJOUhU https://trialofcode.org/
