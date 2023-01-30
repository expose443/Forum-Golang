package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/pkg"
)

func (app *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		pkg.RenderTemplate(w, "signin.html", Messages)
		pkg.ClearStruct(&Messages)

	case http.MethodPost:

		user, err := getUser(r)
		if err != nil {
			Messages.Message = "Wrong password or email"
			http.Redirect(w, r, "/sign-in", http.StatusFound)
			log.Println(err)
			return
		}

		session, err := app.authService.Login(&user)
		if err != nil {
			Messages.Message = "Wrong password or email"
			http.Redirect(w, r, "/sign-in", http.StatusFound)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   session.Token,
			Expires: session.Expiry,
		})

		Sessions = append(Sessions, session)

		http.Redirect(w, r, "/", http.StatusFound)

	default:
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *App) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		pkg.RenderTemplate(w, "signup.html", Messages)
		pkg.ClearStruct(&Messages)

	case http.MethodPost:
		user, err := getUser(r)
		if err != nil {
			Messages.Message = "Wrong user data"
			http.Redirect(w, r, "/sign-up", http.StatusFound)
			log.Println(err)
			return
		}
		err = app.authService.Register(&user)
		if err != nil {
			log.Printf("user %s sign up was failed\n", user.Email)
			Messages.Message = "user exist"
			http.Redirect(w, r, "/sign-up", http.StatusFound)
			return
		}
		http.Redirect(w, r, "/sign-in", http.StatusFound)
		return
	default:
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
	}
}

func (app *App) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	c, err := r.Cookie("session_token")
	if err == nil {
		app.authService.Logout(c.Value)
	}
	http.Redirect(w, r, "/welcome", http.StatusFound)
}

func getUser(r *http.Request) (model.User, error) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	nameRegex, err := regexp.Compile("[a-zA-Z0-9_-]{3,16}")
	if err != nil {
		return model.User{}, errors.New("name regex fail")
	}

	emailRegex, err := regexp.Compile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,6}`)
	if err != nil {
		return model.User{}, errors.New("name regex fail")
	}

	passwordRegex, err := regexp.Compile("[a-zA-Z0-9!@#$%^&*()_+=-]{8,}")
	if err != nil {
		return model.User{}, errors.New("pass regex fail")
	}

	usernameIsValid := nameRegex.MatchString(username)
	emailIsValid := emailRegex.MatchString(email)
	passwordIsValid := passwordRegex.MatchString(password)

	switch r.URL.Path {
	case "/sign-in":
		if emailIsValid && passwordIsValid {
			return model.User{
				Email:    email,
				Password: password,
			}, nil
		} else {
			return model.User{}, errors.New("invalid user data for sign in")
		}
	case "/sign-up":
		if passwordIsValid && usernameIsValid && emailIsValid {
			return model.User{
				Username: username,
				Email:    email,
				Password: password,
			}, nil
		} else {
			return model.User{}, errors.New("invalid user data for sign up")
		}
	default:
		return model.User{}, fmt.Errorf("this url path was not found %s", r.URL.Path)
	}
}

var Sessions []model.Session

func (app *App) ClearSession() {
	all, err := app.sessionService.GetAllSessionsTime()
	if err != nil {
		fmt.Println("error when get all session time", err.Error())
	}
	Sessions = all
	for {
		time.Sleep(time.Second)
		for i, v := range Sessions {
			if v.Expiry.Before(time.Now()) {
				err := app.sessionService.DeleteSession(v.Token)
				if i == len(Sessions)-1 {
					Sessions = Sessions[:i]
				} else {
					Sessions = append(Sessions[:i], Sessions[i+1:]...)
				}
				if err != nil {
					fmt.Println("session delete was fail", err.Error())
				} else {
					fmt.Printf("session for %s was delete\n", v.Username)
				}
			}
		}
	}
}
