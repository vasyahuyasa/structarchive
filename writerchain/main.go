package main

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/rand"
	"os"
)

const mockSize = 100000

type mock struct {
	IntData []int
	IntMap  map[int]int
}

func main() {
	f, err := os.Create("test.json.gz.b64")
	if err != nil {
		log.Fatal("can not create test file: ", err)
	}
	defer f.Close()

	m := makeMock()
	err = encode(f, m)
	if err != nil {
		log.Fatal("can not encode data to file: ", err)
	}
}

// makeMock generate large amount of mock data
func makeMock() *mock {
	m := &mock{
		IntData: make([]int, 0, mockSize),
		IntMap:  make(map[int]int, mockSize),
	}

	for i := 0; i < mockSize; i++ {
		m.IntData = append(m.IntData, rand.Int())
		m.IntMap[i] = rand.Int()
	}

	return m
}

func encode(w io.Writer, data interface{}) error {
	encB64 := base64.NewEncoder(base64.StdEncoding, w)
	encGz := gzip.NewWriter(encB64)
	//encFl, err := flate.NewWriter(w, flate.BestCompression)
	//if err != nil {
	//	return errors.New("can not create zlib writer: " + err.Error())
	//}
	//encLz := lzw.NewWriter(w, lzw.LSB, 8)

	/*
		encZ, err := zlib.NewWriterLevel(w, flate.BestCompression)
		if err != nil {
			return errors.New("can not create zlib writer: " + err.Error())
		}
	*/

	encJson := json.NewEncoder(encGz)
	err := encJson.Encode(data)
	if err != nil {
		return errors.New("can not encode JSON: " + err.Error())
	}

	/*
		err = encLz.Close()
		if err != nil {
			return errors.New("can not close lzw writer: " + err.Error())
		}
	*/

	err = encB64.Close()
	if err != nil {
		return errors.New("can not close base64 encoder: " + err.Error())
	}

	return nil
}
