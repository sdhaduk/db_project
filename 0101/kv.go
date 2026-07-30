package db0101

type KV struct {
	mem map[string][]byte
}

func (kv *KV) Open() error {
	kv.mem = map[string][]byte{} // empty
	return nil
}

func (kv *KV) Close() error { return nil }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error) {
	str_key := string(key)	
	
	val, ok = kv.mem[str_key]
	
	if !ok {
		return nil, ok, err
	}
	
	return val, ok, err
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error) {
	str_key := string(key)
	kv.mem[str_key] = val
	updated = true
	
	return updated, err
}

func (kv *KV) Del(key []byte) (deleted bool, err error) {
	str_key := string(key)

	_, ok := kv.mem[str_key]
	if !ok {
		return deleted, err
	}

	delete(kv.mem, str_key)
	deleted = true

	return deleted, err
}

// QzBQWVJJOUhU https://trialofcode.org/
