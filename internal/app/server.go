package app

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/with-insomnia/Forum-Golang/internal/config"
)

func (app *App) Run(cfg config.Http) *http.Server {
	authPaths := []string{
		"/",
		"/reaction",
		"/comment",
		"/filter",
		"/post",
		"/logout",
		"/welcome",
		"/sign-in",
		"/sign-up",
		"/welcome/filter",
		"/welcome/comment",
		"/post/like",
		"/post/dislike",
		"/comment/dislike",
		"/comment/like",
	}

	AddAuthPaths(authPaths...)

	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	mux.HandleFunc("/", app.authorizedMiddleware(app.IndexHandler))
	mux.HandleFunc("/post", app.authorizedMiddleware(app.PostHandler))
	mux.HandleFunc("/comment", app.authorizedMiddleware(app.CommentHandler))
	mux.HandleFunc("/post/like", app.authorizedMiddleware(app.ReactionHandler))
	mux.HandleFunc("/post/dislike", app.authorizedMiddleware(app.ReactionHandler))
	mux.HandleFunc("/comment/like", app.authorizedMiddleware(app.ReactionHandler))
	mux.HandleFunc("/comment/dislike", app.authorizedMiddleware(app.ReactionHandler))
	mux.HandleFunc("/filter", app.authorizedMiddleware(app.FilterHandler))
	mux.HandleFunc("/logout", app.LogoutHandler)

	mux.HandleFunc("/welcome/filter", app.nonAuthorizedMiddleware(app.FilterWelcomeHandler))
	mux.HandleFunc("/welcome", app.nonAuthorizedMiddleware(app.WelcomeHandler))
	mux.HandleFunc("/sign-in", app.nonAuthorizedMiddleware(app.LoginHandler))
	mux.HandleFunc("/sign-up", app.nonAuthorizedMiddleware(app.RegisterHandler))
	mux.HandleFunc("/welcome/comment", app.nonAuthorizedMiddleware(app.CommentWelcomeHandler))

	fs := http.FileServer(http.Dir("./templates/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	server := &http.Server{
		ReadTimeout:  time.Second * time.Duration(cfg.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(cfg.WriteTimeout),
		IdleTimeout:  time.Second * time.Duration(cfg.IdleTimeout),
		Addr:         cfg.Port,
		Handler:      mux,
	}
	return server
}
