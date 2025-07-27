// cmd/nanospatial/main.go
package main

import (
	"flag"
	"log"
	"os"

	"nanospatial/internal/nanospatial"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := nanospatial.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
