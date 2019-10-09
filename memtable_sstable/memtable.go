package main

import (
	"sync"
)

type value struct {
	tombstone bool
	data      []byte
}

type Memtable struct {
	mu   *sync.RWMutex
	data map[string]value
}

func NewMemtable() *Memtable {
	return &Memtable{
		mu:   &sync.RWMutex{},
		data: map[string]value{},
	}
}

// Set value to given key in memtable
func (t *Memtable) Set(key string, data []byte) {
	t.mu.Lock()
	t.data[key] = value{
		data: data,
	}
	t.mu.Unlock()
}

// Delete key from memtable, if key exist then record market with tombstone otherwise do nothing
func (t *Memtable) Delete(key string) {
	t.mu.Lock()
	t.data[key] = value{
		tombstone: true,
	}
	t.mu.Unlock()
}

// Get find key in memtable, if it is not found second return arg is false
func (t *Memtable) Get(key string) ([]byte, bool) {
	var val []byte
	var isOK bool
	t.mu.RLock()
	v, ok := t.data[key]
	if ok && !v.tombstone {
		isOK = true
		val = v.data
	}
	t.mu.RUnlock()

	return val, isOK
}
