package db0202

import "fmt"

type Schema struct {
	Table string
	Cols  []Column
	PKey  []int // indexes of primary key columns
}

type Column struct {
	Name string
	Type CellType
}

type Row []Cell

func (schema *Schema) NewRow() Row {
	return make(Row, len(schema.Cols))
}

func (row Row) EncodeKey(schema *Schema) (key []byte) {
	key = append([]byte(schema.Table), 0x00)

	for _, key_idx := range schema.PKey {
		key = row[key_idx].Encode(key)
	}

	return key
}

func (row Row) EncodeVal(schema *Schema) (val []byte) {
	set := make(map[int]struct{}, len(schema.PKey))	
	for _, key_idx := range schema.PKey {
		set[key_idx] = struct{}{}
	}

	for val_idx, cell := range row {
		_, ok := set[val_idx]
		if !ok {
			val = cell.Encode(val)
		}
	}

	return val
}

func (row Row) DecodeKey(schema *Schema, key []byte) (err error) {
	table_name := string(key[:len(schema.Table)])
	if table_name != schema.Table {
		return fmt.Errorf("table names do not match")
	}

	key = key[len(schema.Table) + 1:]

	for _, key_idx := range schema.PKey {
		row[key_idx].Type = schema.Cols[key_idx].Type
		key, err = row[key_idx].Decode(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func (row Row) DecodeVal(schema *Schema, val []byte) (err error) {
	set := make(map[int]struct{}, len(schema.PKey))	
	for _, key_idx := range schema.PKey {
		set[key_idx] = struct{}{}
	}

	for val_idx, _ := range row {
		_, ok := set[val_idx]
		if !ok {
			row[val_idx].Type = schema.Cols[val_idx].Type
			val, err = row[val_idx].Decode(val)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// QzBQWVJJOUhU https://trialofcode.org/
