package app

import (
	"github.com/with-insomnia/Forum-Golang/internal/config"
	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/internal/service"
)

var Messages model.Data

type App struct {
	authService    service.AuthService
	sessionService service.SessionService
	postService    service.PostService
	userService    service.UserService
	cfg            config.Config
}

func NewAppService(
	authService service.AuthService,
	sessionService service.SessionService,
	postService service.PostService,
	userService service.UserService,
	cfg config.Config,
) App {
	return App{
		authService:    authService,
		sessionService: sessionService,
		postService:    postService,
		userService:    userService,
		cfg:            cfg,
	}
}
