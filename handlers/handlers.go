package handlers

import (
	"net/http"

	"github.com/Zigelzi/taxi-pudasjarvi/components"
)

func Index(w http.ResponseWriter, r *http.Request) {
	component := components.Index()
	component.Render(r.Context(), w)
}

func Services(w http.ResponseWriter, r *http.Request) {
	component := components.Services()
	component.Render(r.Context(), w)
}

func Prices(w http.ResponseWriter, r *http.Request) {
	component := components.Prices()
	component.Render(r.Context(), w)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	component := components.Contact()
	component.Render(r.Context(), w)
}
