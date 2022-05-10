package kv

type (
	Store struct {
	}
)

func Open(path string) (*Store, error) {
	// Open path, load and verifydata, replay pending transactions etc.
	return nil, nil
}

// Put persists (key, value) to the store.
func (s *Store) Put(key string, value interface{}) error { // ...
	return nil
}

// Get looks up the value associated with key.
func (s *Store) Get(key string) (interface{}, error) { // ...
	return nil, nil
}

// Close waits for any pending transactions to complete and then
// cleanly shuts down the KV store.
func (s *Store) Close() error { // ...
	return nil
}
