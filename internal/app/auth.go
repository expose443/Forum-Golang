package app

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/google/uuid"
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
			pkg.ErrorHandler(w, http.StatusBadRequest)
			log.Println(err)
			return
		}
		sessionDb, err := app.sessionService.GetSessionByUserID(int(user.ID))
		if err != nil {
			log.Printf("session for user_id %d is not found\n", user.ID)
		}
		if sessionDb.UserId == user.ID {
			err := app.sessionService.DeleteSession(sessionDb.Token)
			if err != nil {
				log.Println(err)
			}
		}
		sessionToken := uuid.NewString()
		expiry := time.Now().Add(10 * time.Minute)
		session := model.Session{
			UserId: user.ID,
			Token:  sessionToken,
			Expiry: expiry,
		}
		_, err = app.authService.Login(user, &session)
		if err != nil {
			log.Printf("user %s sign in was failed\n", user.Email)
			Messages.Message = "incorrect data"
			http.Redirect(w, r, "/sign-in", http.StatusFound)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   sessionToken,
			Expires: expiry,
		})

		log.Printf("user %s sign in was successfully\n", user.Email)
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
			pkg.ErrorHandler(w, http.StatusBadRequest)
			log.Println(err)
			return
		}
		err = app.authService.Register(user)
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
	if err != nil {
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = app.authService.Logout(c.Value)
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/welcome", http.StatusFound)
}

func getUser(r *http.Request) (*model.User, error) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	nameRegex, _ := regexp.Compile("[a-zA-z]{6,30}")
	emailRegex, _ := regexp.Compile("^[a-zA-Z0-9_\\-\\.]+@[a-zA-Z0-9\\-\\.]+\\.[a-zA-Z]{2,}$")
	passwordRegex, _ := regexp.Compile("[a-zA-Z0-9]{6, 30}")

	usernameIsValid := nameRegex.MatchString(username)
	emailIsValid := emailRegex.MatchString(email)
	passwordIsValid := passwordRegex.MatchString(password)

	switch r.URL.Path {
	case "/sign-in":
		if emailIsValid && passwordIsValid {
			return &model.User{
				Email:    email,
				Password: password,
			}, nil
		}
	case "/sign-up":
		if emailIsValid && passwordIsValid && usernameIsValid {
			return &model.User{
				Username: username,
				Email:    email,
				Password: password,
			}, nil
		}
	}
	return nil, errors.New("invalid user data")
}
