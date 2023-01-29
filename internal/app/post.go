package app

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/pkg"
)

func (app *App) PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		pkg.RenderTemplate(w, "createpost.html", model.Data{})
		return
	case http.MethodPost:
		r.ParseForm()
		title := r.FormValue("title")
		message := r.FormValue("message")
		genre := r.Form["category"]

		user, ok := r.Context().Value(keyUserType(keyUser)).(model.User)
		if !ok {
			pkg.ErrorHandler(w, http.StatusUnauthorized)
			return
		}

		post := model.Post{
			Title:      title,
			Content:    message,
			Category:   strings.Join(genre, " "),
			Author:     user,
			CreateTime: time.Now().Format(time.RFC822),
		}

		status, err := app.postService.CreatePost(&post)
		if err != nil {
			log.Println(err)
			switch status {
			case http.StatusInternalServerError:
				pkg.ErrorHandler(w, http.StatusInternalServerError)
				return
			case http.StatusBadRequest:
				pkg.ErrorHandler(w, http.StatusBadRequest)
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusFound)

	default:
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}
