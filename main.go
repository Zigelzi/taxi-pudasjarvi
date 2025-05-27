package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	const port = ":3000"
	component := index("Testing")
	http.Handle("/", templ.Handler(component))
	fmt.Println("Starting server on port: ", port)
	http.ListenAndServe(port, nil)
}
