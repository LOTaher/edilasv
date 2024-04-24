package main

import (
	"log"

	"github.com/LOTaher/softbase"
	"github.com/LOTaher/softbase/cmd"
    "github.com/LOTaher/softbase/core"
)

func main() {
	app := softbase.New()
    db := core.NewStore(2)
	// Serve Command
    serveCmd := cmd.Serve(db)

	if err := app.Start(serveCmd); err != nil {
		log.Fatal(err)
	}
}
