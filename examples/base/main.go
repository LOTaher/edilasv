package main

import (
	"log"

	"github.com/LOTaher/softbase"
	"github.com/LOTaher/softbase/cmd"
    "github.com/LOTaher/softbase/core"
)

func main() {
    var db *core.Store 

    if softbase.HasDatabase() {
        db = softbase.LoadDatabase()
    } else {
        db = core.NewStore(2)
    }

	app := softbase.New(db)
    
	// Serve Command
    serveCmd := cmd.Serve(db)

	if err := app.Start(serveCmd); err != nil {
		log.Fatal(err)
	}
}
