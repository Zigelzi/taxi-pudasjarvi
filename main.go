package main

import (
	"fmt"
	"net/http"

	"github.com/Zigelzi/taxi-pudasjarvi/handlers"
)

func main() {
	const port = ":3000"
	// Static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Routes
	http.HandleFunc("/", handlers.Index)
	fmt.Println("Starting server on port: ", port)
	http.ListenAndServe(port, nil)
}
