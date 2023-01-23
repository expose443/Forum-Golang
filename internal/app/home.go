package app

import (
	"net/http"

	"github.com/with-insomnia/Forum-Golang/pkg"
)

func (app *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	switch r.URL.Path {
	case "/":
		// index
	case "/welcome":
		// unauth
	default:
		pkg.ErrorHandler(w, http.StatusNotFound)
		return
	}
}
