package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Zigelzi/taxi-pudasjarvi/handlers"
)

func main() {
	// Static files
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))

	// Routes
	http.HandleFunc("GET /{$}", handlers.Index)
	http.HandleFunc("GET /palvelut", handlers.Services)
	http.HandleFunc("GET /hinnasto", handlers.Prices)
	http.HandleFunc("GET /yhteystiedot", handlers.Contact)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Starting server on port:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Error starting server: ", err)
	}
}
