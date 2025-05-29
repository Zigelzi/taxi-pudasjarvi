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
	http.HandleFunc("GET /{$}", handlers.Index)
	http.HandleFunc("GET /palvelut", handlers.Services)
	http.HandleFunc("GET /hinnasto", handlers.Prices)
	http.HandleFunc("GET /yhteystiedot", handlers.Contact)
	fmt.Println("Starting server on port: ", port)
	http.ListenAndServe(port, nil)
}
