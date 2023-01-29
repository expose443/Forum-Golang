package app

import (
	"net/http"

	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/pkg"
)

func (app *App) FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	category := r.URL.Query().Get("category")
	user, ok := r.Context().Value(keyUserType(keyUser)).(model.User)
	if !ok {
		pkg.ErrorHandler(w, http.StatusUnauthorized)
		return
	}

	data, status := app.postService.GetFilterPosts(category, user)
	switch status {
	case http.StatusInternalServerError:
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	case http.StatusBadRequest:
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	case http.StatusOK:
		pkg.RenderTemplate(w, "filter.html", data)
	}
}

func (app *App) FilterWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	category := r.URL.Query().Get("category")

	data, status := app.postService.GetWelcomeFilterPosts(category)
	switch status {
	case http.StatusInternalServerError:
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	case http.StatusBadRequest:
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	case http.StatusOK:
		pkg.RenderTemplate(w, "welcome.html", data)
	}
}
