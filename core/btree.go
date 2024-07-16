package core

import (
	"sync"

	"github.com/google/btree"
)

type Item struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func (i Item) Less(than btree.Item) bool {
	return i.Key < than.(Item).Key
}

type Store struct {
	tree *btree.BTree
	mu   sync.RWMutex
}

func NewStore(degree int) *Store {
	return &Store{
		tree: btree.New(degree),
	}
}

func (s *Store) Insert(Key string, Value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item := Item{Key: Key, Value: Value}
	s.tree.ReplaceOrInsert(item)
}

func (s *Store) Get(Key string) (Value interface{}, ok bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item := s.tree.Get(Item{Key: Key})
	if item != nil {
		return item.(Item).Value, true
	}

	return nil, false
}

func (s *Store) Delete(Key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tree.Delete(Item{Key: Key})
}

func (s *Store) Update(Key string, Value interface{}) {
	s.Delete(Key)
	s.Insert(Key, Value)
}

func (s *Store) Has(Key string) bool {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return s.tree.Has(Item{Key: Key})
}

func (s *Store) GetAll() []Item {
    s.mu.RLock()
    defer s.mu.RUnlock()
    var items []Item
    s.tree.Ascend(func(i btree.Item) bool {
        items = append(items, i.(Item))
        return true
    })
    return items
}
