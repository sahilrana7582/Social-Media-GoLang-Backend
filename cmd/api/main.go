package main

import (
	"fmt"
	"log"
)

func main() {
	config := config{
		addr: ":4000",
	}

	app := &application{
		config: config,
	}

	fmt.Println("Starting server on port", config.addr)

	log.Fatal(app.run())
}
