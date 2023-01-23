package app

import (
	"net/http"

	"github.com/with-insomnia/Forum-Golang/pkg"
)

func (app *App) PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// TODO get func
	case http.MethodPost:
		// TODO post func
	default:
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}
