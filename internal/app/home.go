package app

import (
	"net/http"

	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/pkg"
)

func (app *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	switch r.URL.Path {
	case "/":
		pkg.RenderTemplate(w, "index.html", model.Data{})
	default:
		pkg.ErrorHandler(w, http.StatusNotFound)
		return
	}
}

func (app *App) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	pkg.RenderTemplate(w, "welcome.html", model.Data{})
}
