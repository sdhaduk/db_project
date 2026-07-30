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
		
		kv.mem[string(ent.key)] = ent.val
	}

	return nil
}

func (kv *KV) Close() error { return kv.log.Close() }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error) {
	val, ok = kv.mem[string(key)]
	return
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error)

func (kv *KV) Del(key []byte) (deleted bool, err error)

// QzBQWVJJOUhU https://trialofcode.org/

