package delivery

import (
	"net/http"

	"github.com/with-insomnia/Forum-Golang/pkg"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// TODO func for get
	case http.MethodPost:
		// TODO func for post
	default:
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// TODO func for get
	case http.MethodPost:
		// TODO func for post
	default:
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	// TODO Logout logic
}
