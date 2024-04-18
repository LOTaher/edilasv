package storage

import (

	"github.com/LOTaher/softbase/pkg/btree" 
)

type Database struct {
	Store *btree.Store
}

func NewDatabase(degree int) *Database {
	store := btree.NewStore(degree) 
	return &Database{
		Store: store,
	}
}

func (db *Database) AddItem(key string, value interface{}) {
	db.Store.Insert(key, value)
}

func (db *Database) GetItem(key string) (interface{}, bool) {
	return db.Store.Get(key)
}

func (db *Database) DeleteItem(key string) {
	db.Store.Delete(key)
}

func (db *Database) UpdateItem(key string, value interface{}) {
	db.Store.Update(key, value)
}

// TODO
// Serialize and Deserialize database so it persists across multiple sessions


