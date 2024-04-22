package main

import (
    "log"

	"github.com/LOTaher/softbase"
)

func main() {
	app := softbase.New()

	// Optional Plugin Flags:

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
