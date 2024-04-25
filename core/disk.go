package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"os"

	"github.com/google/btree"
)

func SerializeNode(item *Item) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(item)
	if err != nil {
		fmt.Printf("SerializeNode failed: %v", err)
	}

	return buf.Bytes()

}

func DeserializeNode(data []byte) *Item {
	buf := bytes.NewBuffer(data)
	var item Item
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&item)
	if err != nil {
		fmt.Printf("DeserializeNode failed: %v", err)
	}

	return &item
}

func serializeItemIterator(enc *gob.Encoder) btree.ItemIterator {
	return func(item btree.Item) bool {
		if err := enc.Encode(item); err != nil {
			fmt.Printf("Failed to encode item: %v", err)
			return false
		}
		return true
	}
}

func SaveToDisk(filename string, tree *btree.BTree) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	tree.Ascend(serializeItemIterator(enc))

	return nil
}

func LoadFromDisk(filename string, tree *btree.BTree) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	dec := gob.NewDecoder(file)
	var item Item
	for {
		err := dec.Decode(&item)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		tree.ReplaceOrInsert(item)
	}
	return nil
}
