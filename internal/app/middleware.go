package app

import (
	"context"
	"net/http"
	"time"

	"github.com/with-insomnia/Forum-Golang/pkg"
)

var AuthPaths = make(map[string]struct{})

func AddAuthPaths(paths ...string) {
	for _, path := range paths {
		AuthPaths[path] = struct{}{}
	}
}

func (app *App) authorizedMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := AuthPaths[r.URL.Path]; !ok {
			pkg.ErrorHandler(w, http.StatusNotFound)
			return
		}
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "/welcome", http.StatusFound)
			return
		}

		session, err := app.sessionService.GetSessionByToken(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/welcome", http.StatusFound)
			return
		}
		if session.Expiry.Before(time.Now()) {
			http.Redirect(w, r, "/welcome", http.StatusFound)
			return
		}

		user, err := app.userService.GetUserByToken(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/welcome", http.StatusFound)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *App) nonAuthorizedMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := AuthPaths[r.URL.Path]; !ok {
			pkg.ErrorHandler(w, http.StatusNotFound)
			return
		}
		c, err := r.Cookie("session_token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		checkSessionFromDb, err := app.sessionService.GetSessionByToken(c.Value)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		if checkSessionFromDb.Expiry.Before(time.Now()) {
			next.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	})
}
