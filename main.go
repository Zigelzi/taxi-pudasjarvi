package main

import (
	"os"

	filegeneration "github.com/Zigelzi/taxi-pudasjarvi/file_generation"
	"github.com/Zigelzi/taxi-pudasjarvi/server"
	"github.com/Zigelzi/taxi-pudasjarvi/views"
)

func main() {
	var arg string
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	if arg == "static" {
		appViews := views.Get()
		for route, component := range appViews {
			filegeneration.CreateStaticFiles(route, component)

		}
		return
	}
	port := os.Getenv("PORT")

	serv := server.Server{Port: port}
	serv.Start(true)
}
