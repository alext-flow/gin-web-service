package main

import (
	"api/server"
	"log"
)

func main() {
	app := server.App{}
	err := app.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	app.ConfigureRoutes()
	app.Serve()
}
