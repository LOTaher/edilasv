package core

import (
    "os"
    "encoding/gob"
    "fmt"
    "io"

    "github.com/google/btree"
)

func init() {
    gob.Register(Item{})
    gob.Register(map[string]interface{}{})
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

func (s *Store) SaveToDisk(filename string) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    enc := gob.NewEncoder(file)

    s.tree.Ascend(serializeItemIterator(enc))
    return err
}

// Custom deserialization function
func (s *Store) LoadFromDisk(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    dec := gob.NewDecoder(file)
    var item Item
    for {
        err = dec.Decode(&item)
        if err != nil {
            if err == io.EOF {
                break
            }
            return err
        }
        s.tree.ReplaceOrInsert(item)
    }
    return nil
}
