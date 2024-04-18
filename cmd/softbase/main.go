package main

import (
	"fmt"

    "github.com/fatih/color"
)

func main() {
    fmt.Print("∘˙○˚.•\n")
    started := color.New(color.FgCyan, color.Bold)
    started.Printf("> Server started at %s\n", "http://localhost:8080")
    fmt.Printf("  - REST API: %s\n", "http://localhost:8080/api/")
    fmt.Printf("  - Admin Panel: %s\n", "http://localhost:8080/")
    fmt.Println("∘˙○˚.•")
}
