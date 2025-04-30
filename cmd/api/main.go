package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sahilrana7582/Social-Media-GoLang-Backend/pkg/db"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	db.InitDB()

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
	log.Println("Server started on port", config.addr)
}
