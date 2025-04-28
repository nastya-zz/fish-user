package event

import (
	"user/internal/service"
)

type Processor struct {
	userService service.UserService
}

func New(userService service.UserService) service.EventsService {
	return &Processor{userService: userService}
}
