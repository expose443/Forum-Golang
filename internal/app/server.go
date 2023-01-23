package app

import (
	"net/http"
	"os"
	"time"
)

func (app *App) Run() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.IndexHandler)
	mux.HandleFunc("/sign-in", app.LoginHandler)
	mux.HandleFunc("/sign-up", app.RegisterHandler)
	mux.HandleFunc("/post", app.PostHandler)
	mux.HandleFunc("/comment", app.CommentHandler)
	mux.HandleFunc("/like", app.ReactionHandler)
	mux.HandleFunc("/filter", app.FilterHandler)
	mux.HandleFunc("/logout", app.LogoutHandler)

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
