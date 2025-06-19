package handlers

import (
	"net/http"

	"github.com/Zigelzi/taxi-pudasjarvi/components"
)

func Index(w http.ResponseWriter, r *http.Request) {
	component := components.Index()
	component.Render(r.Context(), w)
}
