package main

import (
	"log"
)

func main() {
	m := NewMemtable()
	m.Set("test", []byte("hello memtable"))
	data, ok := m.Get("test")
	if ok {
		log.Println("get:", string(data))
	} else {
		log.Println("value for key test not found")
	}

	m.Delete("test")
	data, ok = m.Get("test")
	if ok {
		log.Println("get:", string(data))
	} else {
		log.Println("value for key test not found")
	}
}
