package app

import (
	"net/http"
	"strconv"

	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/pkg"
)

func (app *App) ReactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Query().Get("path")
	if path == "" {
		path = "/"
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	user, ok := r.Context().Value(keyUserType(keyUser)).(model.User)
	if !ok {
		pkg.ErrorHandler(w, http.StatusUnauthorized)
		return
	}
	switch r.URL.Path {
	case "/post/like":
		status := app.postService.LikePost(id, int(user.ID))
		switch status {
		case http.StatusInternalServerError:
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		case http.StatusBadRequest:
			pkg.ErrorHandler(w, http.StatusBadRequest)
			return
		case http.StatusOK:
			http.Redirect(w, r, path, http.StatusFound)

		}
	case "/post/dislike":
		status := app.postService.DislikePost(id, int(user.ID))
		switch status {
		case http.StatusInternalServerError:
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		case http.StatusBadRequest:
			pkg.ErrorHandler(w, http.StatusBadRequest)
			return
		case http.StatusOK:
			http.Redirect(w, r, path, http.StatusFound)
		}
	case "/comment/like":
		status := app.postService.LikeComment(id, int(user.ID))
		switch status {
		case http.StatusInternalServerError:
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		case http.StatusBadRequest:
			pkg.ErrorHandler(w, http.StatusBadRequest)
			return
		case http.StatusOK:
			http.Redirect(w, r, path, http.StatusFound)
		}
	case "/comment/dislike":
		status := app.postService.DisLikeComment(id, int(user.ID))
		switch status {
		case http.StatusInternalServerError:
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		case http.StatusBadRequest:
			pkg.ErrorHandler(w, http.StatusBadRequest)
			return
		case http.StatusOK:
			http.Redirect(w, r, path, http.StatusFound)
		}
	default:
		pkg.ErrorHandler(w, http.StatusNotFound)
		return
	}
}
