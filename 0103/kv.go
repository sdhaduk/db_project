package db0103

type KV struct {
	log Log
	mem map[string][]byte
}

func (kv *KV) Open() error {
	if err := kv.log.Open(); err != nil {
		return err	
	}
	kv.mem = map[string][]byte{}
	ent := &Entry{}

	for {
		eof, err := kv.log.Read(ent)

		if err != nil {
			return err
		}
		if eof {
			break
		}

		if ent.deleted {
			delete(kv.mem, string(ent.key))
		} else {
			kv.mem[string(ent.key)] = ent.val
		}
	}

	return nil
}

func (kv *KV) Close() error { return kv.log.Close() }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error) {
	val, ok = kv.mem[string(key)]
	return
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error) {	
	ent := &Entry{key: key, val: val,}
	err = kv.log.Write(ent)
	if err != nil {
		return updated, err
	}
	kv.mem[string(key)] = val

	updated = true
	return updated, err
}

func (kv *KV) Del(key []byte) (deleted bool, err error) {
	_, ok, err := kv.Get(key)
	if err != nil {
		return deleted, err
	}
	if !ok {
		return deleted, err
	}
	
	ent := &Entry{key: key, deleted: true}
	err = kv.log.Write(ent)
	if err != nil {
		return deleted, err
	}
	delete(kv.mem, string(key))
	deleted = true
	
	return deleted, nil
}

// QzBQWVJJOUhU https://trialofcode.org/

