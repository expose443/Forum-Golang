package app

import (
	"net/http"
	"os"
	"time"

	delivery "github.com/with-insomnia/Forum-Golang/internal/delivery/http"
)

func (app *App) Run() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", delivery.IndexHandler)
	mux.HandleFunc("/sign-in", delivery.LoginHandler)
	mux.HandleFunc("/sign-up", delivery.RegisterHandler)
	mux.HandleFunc("/post", delivery.PostHandler)
	mux.HandleFunc("/comment", delivery.CommentHandler)
	mux.HandleFunc("/like", delivery.ReactionHandler)
	mux.HandleFunc("/filter", delivery.FilterHandler)
	mux.HandleFunc("/logout", delivery.LogoutHandler)

	fs := http.FileServer(http.Dir("./templates/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	port, ok := os.LookupEnv("FORUM_PORT")
	if !ok {
		port = "8080"
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
