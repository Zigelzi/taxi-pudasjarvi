package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/Zigelzi/taxi-pudasjarvi/handlers"
)

type Server struct {
	Port string
}

func (s Server) Start(isDev bool) {
	// Static files
	http.Handle("/assets/",
		disableCacheInDevMode(
			http.StripPrefix("/assets",
				http.FileServer(http.Dir("assets"))),
			isDev))

	// Routes
	http.HandleFunc("GET /{$}", handlers.Index)

	if s.Port == "" {
		s.Port = "8080"
	}
	log.Println("Starting server on port:", s.Port)
	err := http.ListenAndServe(":"+s.Port, nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Error starting server: ", err)
	}
}

func disableCacheInDevMode(next http.Handler, isDev bool) http.Handler {
	if !isDev {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
