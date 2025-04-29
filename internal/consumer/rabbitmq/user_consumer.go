package rabbitmq

import (
	"bytes"
	"context"
	"encoding/gob"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"user/internal/consumer"
	"user/internal/model"
	"user/internal/service"
)

const (
	queueName = "user"
)

type UserConsumer struct {
	ch        *amqp.Channel
	processor service.EventsService
}

func NewUserConsumer(ch *amqp.Channel, processor service.EventsService) consumer.Consumer {
	return &UserConsumer{
		ch:        ch,
		processor: processor,
	}
}

// TODO: решить вопрос с обработкой ошибок, паники
func (u UserConsumer) Start(ctx context.Context) {
	messages, err := u.ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("stopping event consumer by ctx")
				return
			case message := <-messages:
				log.Printf("Message: %s\n", message.Body)
				e := event(message.Body)
				err = u.processor.Process(ctx, *e)
				if err != nil || e == nil {
					log.Printf("err: %s\n", err)
					log.Printf("err: %s\n", e)
				}
			case <-sigchan:
				log.Println("Interrupt detected!")
				os.Exit(0)
			}
		}
	}()
}

func event(msg []byte) *model.Event {
	var e model.Event
	buf := bytes.NewBuffer(msg)

	if err := gob.NewDecoder(buf).Decode(&e); err != nil {
		return nil
	}

	log.Println("event:", e)

	return &e
}
