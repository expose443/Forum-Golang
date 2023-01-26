package app

import (
	"net/http"

	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/pkg"
)

func (app *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		pkg.ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	user, ok := r.Context().Value(keyUser).(model.User)
	if !ok {
		pkg.ErrorHandler(w, http.StatusUnauthorized)
		return
	}

	post, err := app.postService.GetAllPosts()
	if err != nil {
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	data := model.Data{
		Posts: post,
		User:  user,
		Genre: "/",
	}
	pkg.RenderTemplate(w, "index.html", data)
}

func (app *App) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	pkg.RenderTemplate(w, "welcome.html", model.Data{})
}
