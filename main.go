package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Zigelzi/taxi-pudasjarvi/components"
	filegeneration "github.com/Zigelzi/taxi-pudasjarvi/file_generation"
	"github.com/Zigelzi/taxi-pudasjarvi/handlers"
)

func main() {

	var arg string
	if len(os.Args) > 0 {
		arg = os.Args[1]
	}

	if arg == "static" {
		createStaticFiles()
		return
	}
	// Static files
	http.Handle("/assets/",
		disableCacheInDevMode(
			http.StripPrefix("/assets",
				http.FileServer(http.Dir("assets")))))

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

var dev = true

func disableCacheInDevMode(next http.Handler) http.Handler {
	if !dev {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func createStaticFiles() {
	rootPath := "public"
	err := os.Mkdir(rootPath, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("failed to create output directory %v", err)
	}
	filePath := filepath.Join(rootPath, "index.html")

	f, err := os.Create(filePath)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("failed to create file %v", err)
	}

	err = components.Index().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to render component: %v", err)
	}

	assetsPath := filepath.Join(rootPath, "assets")
	err = os.Mkdir(assetsPath, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("failed to create assets output directory %v", err)
	}

	tailwindPath := filepath.Join(assetsPath, "tailwind.css")
	err = filegeneration.Copy("./assets/tailwind.css", tailwindPath)
	if err != nil {
		log.Fatalf("failed to copy asset files: %v", err)
	}
}
