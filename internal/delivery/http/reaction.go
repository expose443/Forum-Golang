package delivery

import (
	"net/http"

	"github.com/with-insomnia/Forum-Golang/pkg"
)

func ReactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	switch r.URL.Path {
	case "/post/like-post":
		// TODO logic
	case "/post/dislike-post":
		// TODO logic
	case "/comment/like-comment":
		// TODO logic
	case "/comment/dislike-comment":
		// TODO logic
	default:
		pkg.ErrorHandler(w, http.StatusNotFound)
		return
	}
}
