package main

import (
	"fmt"

    "github.com/google/btree"
)

type Item struct {
    key string
    value interface{}
}

func (i Item) Less(than btree.Item) bool {
    return i.key < than.(Item).key
}

type Store struct {
    tree *btree.BTree
}

func NewStore(degree int) *Store {
    return &Store {
        tree: btree.New(degree),
    }
}

func (s *Store) Put(key string, value interface{}) {
    item := Item {key: key, value: value}
    s.tree.ReplaceOrInsert(item)
}

func (s *Store) Get(key string) (value interface{}, ok bool) {
    item := s.tree.Get(Item{key: key})
    if item != nil {
        return item.(Item).value, true
    }

    return nil, false
}

func (s *Store) Delete(key string) {
    s.tree.Delete(Item{key: key})
}

func main() {
	fmt.Println("init")
}
