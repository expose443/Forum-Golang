package app

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/pkg"
)

func (app *App) CommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Println(err)
			pkg.ErrorHandler(w, http.StatusBadRequest)
			return
		}

		initialPost, status := app.postService.GetAllCommentsAndPostByPostID(int64(id))
		switch status {
		case http.StatusInternalServerError:
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		case http.StatusBadRequest:
			pkg.ErrorHandler(w, http.StatusBadRequest)
			return
		}

		data := model.Data{
			Comment:     initialPost.Comment,
			InitialPost: initialPost,
		}

		pkg.RenderTemplate(w, "commentview.html", data)
	case http.MethodPost:
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Println(err)
			pkg.ErrorHandler(w, http.StatusBadRequest)
			return
		}

		message := r.FormValue("comment")

		path := "/comment?id=" + (r.URL.Query().Get("id"))

		user, ok := r.Context().Value(keyUser).(model.User)
		if !ok {
			pkg.ErrorHandler(w, http.StatusUnauthorized)
			return
		}

		comment := model.Comment{
			PostId:   int64(id),
			Message:  message,
			UserId:   user.ID,
			Username: user.Username,
			Born:     time.Now().Format(time.RFC822),
		}
		status, err := app.postService.CreateComment(&comment)
		log.Println(err)
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
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *App) CommentWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	initialPost, status := app.postService.GetAllCommentsAndPostByPostID(int64(id))
	switch status {
	case http.StatusInternalServerError:
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	case http.StatusBadRequest:
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}

	data := model.Data{
		Comment:     initialPost.Comment,
		InitialPost: initialPost,
	}

	pkg.RenderTemplate(w, "commentunauth.html", data)
}
