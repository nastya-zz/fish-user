package event

import (
	amqp "github.com/rabbitmq/amqp091-go"

	"user/internal/repository"
	"user/internal/service"
)

type Processor struct {
	userService service.UserService
}
type Sender struct {
	eventRepository repository.EventRepository
	ch              *amqp.Channel //TODO: REFACTOR
}

type eventService struct {
	processor Processor
	sender    Sender
}

func New(userService service.UserService, eventRepository repository.EventRepository, ch *amqp.Channel) service.EventsService {
	return &eventService{
		processor: Processor{userService: userService},
		sender:    Sender{eventRepository: eventRepository, ch: ch},
	}
}
