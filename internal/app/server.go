package app

import (
	"net/http"
	"os"
	"time"
)

func (app *App) Run() *http.Server {
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
	}

	AddAuthPaths(authPaths...)

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.authorizedMiddleware(app.IndexHandler))
	mux.HandleFunc("/post", app.authorizedMiddleware(app.PostHandler))
	mux.HandleFunc("/comment", app.authorizedMiddleware(app.CommentHandler))
	mux.HandleFunc("/post/like", app.authorizedMiddleware(app.ReactionHandler))
	mux.HandleFunc("/post/dislike", app.authorizedMiddleware(app.ReactionHandler))
	mux.HandleFunc("/filter", app.authorizedMiddleware(app.FilterHandler))
	mux.HandleFunc("/logout", app.LogoutHandler)

	mux.HandleFunc("/welcome/filter", app.nonAuthorizedMiddleware(app.FilterWelcomeHandler))
	mux.HandleFunc("/welcome", app.nonAuthorizedMiddleware(app.WelcomeHandler))
	mux.HandleFunc("/sign-in", app.nonAuthorizedMiddleware(app.LoginHandler))
	mux.HandleFunc("/sign-up", app.nonAuthorizedMiddleware(app.RegisterHandler))
	mux.HandleFunc("/welcome/comment", app.nonAuthorizedMiddleware(app.CommentWelcomeHandler))

	fs := http.FileServer(http.Dir("./templates/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	port, ok := os.LookupEnv("FORUM_PORT")
	if !ok {
		port = ":8080"
	}

	server := &http.Server{
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 30,
		Addr:         port,
		Handler:      mux,
	}
	return server
}
