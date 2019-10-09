package main

import (
	"bytes"
	"encoding/binary"
	"io"
)

type indexRow struct {
	tombstone bool
	offset    int64
	size      int64
}

type SSTable struct {
	r     io.ReadSeeker
	index map[string]indexRow
}

func NewSSTable(r io.ReadSeeker) *SSTable {
	return &SSTable{
		r:     r,
		index: map[string]indexRow{},
	}
}

func (t *SSTable) LoadIndex(r io.Reader) error {

}

func (t *SSTable) readIndexRow(r io.Reader) (string, indexRow, error) {
	keySizeBuf := make([]byte, 8)
	err := binary.Read(r, binary.LittleEndian, keySizeBuf)
	if err != nil {
		return "", indexRow{}, err
	}

	keySize, err := binary.ReadVarint(bytes.NewBuffer(keySizeBuf))
	if err != nil {
		return "", indexRow{}, err
	}

	key := make([]byte, keySize)
	binary.Read(r, binary.LittleEndian, key)

	var record indexRow
	dataBuf := make([]byte, 1+8+8)
	err = binary.Read(r, binary.LittleEndian, dataBuf)
	if err != nil {
		return "", indexRow{}, err
	}

	record.tombstone = dataBuf[0] != 0

}
