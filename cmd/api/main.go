package main

import (
	"log"
)

func main() {
	config := config{
		addr: ":4000",
	}

	app := &application{
		config: config,
	}

	mux := app.mount()
	if err := app.run(mux); err != nil {
		log.Fatal(err)
	}
}
