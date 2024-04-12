package db

import (

    "github.com/LOTaher/ediasv/internal/btree"
)

type DB struct {
    storage *btree.BTree
}


