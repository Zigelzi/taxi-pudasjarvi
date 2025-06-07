package main

import (
	"os"

	filegeneration "github.com/Zigelzi/taxi-pudasjarvi/file_generation"
	"github.com/Zigelzi/taxi-pudasjarvi/server"
)

func main() {

	var arg string
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	if arg == "static" {
		filegeneration.CreateStaticFiles()
		return
	}
	port := os.Getenv("PORT")

	serv := server.Server{Port: port}
	serv.Start(true)
}
