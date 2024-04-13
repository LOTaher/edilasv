package main

import (
	"fmt"
    
    "github.com/LOTaher/softbase/pkg/btree"
)

func main() {
    fmt.Println("Now serving SoftBase...")
    store := btree.NewStore(50)

    store.Put("key1", "value1")

    value, ok := store.Get("key1")
    if ok {
        fmt.Println("key1:", value)
    }
}
