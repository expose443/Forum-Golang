package app

import (
	"context"
	"net/http"
	"time"
)

func (app *App) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "/sign-in", http.StatusFound)
			return
		}

		session, err := app.sessionService.GetSessionByToken(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/sign-in", http.StatusFound)
			return
		}
		if session.Expiry.Before(time.Now()) {
			http.Redirect(w, r, "/sign-in", http.StatusFound)
			return
		}

		user, err := app.userService.GetUserByToken(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/sign-in", http.StatusFound)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
