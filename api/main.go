package main

import (
	"api/server"
)

func main() {
	app := server.App{}
	app.Initialize()
	app.ConfigureRoutes()
	app.Serve()
}
