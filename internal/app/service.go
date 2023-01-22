package app

import "github.com/with-insomnia/Forum-Golang/internal/service"

type App struct {
	sessionService service.SessionService
	userService    service.UserService
}

func NewAppService(
	sessionService service.SessionService,
) App {
	return App{
		sessionService: sessionService,
	}
}
