package rabbitmq

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"os/signal"
	"syscall"
	"user/internal/consumer"
	"user/internal/model"
	"user/internal/service"
	"user/pkg/logger"
)

const (
	queueName    = "user"
	ExchangeName = "user_events"
)

type UserConsumer struct {
	ch        *amqp.Channel
	processor service.EventsService
}

func NewUserConsumer(ch *amqp.Channel, processor service.EventsService, serviceName string) consumer.Consumer {

	queue, err := ch.QueueDeclare(
		fmt.Sprintf("%s", serviceName),
		true,  // durable
		false, // auto delete
		false, // exclusive
		false, // no wait
		nil,   // args
	)
	if err != nil {
		return nil
	}

	// Привязываем очередь к exchange
	err = ch.QueueBind(
		queue.Name,
		"", // routing key не используется для fanout
		ExchangeName,
		false,
		nil,
	)
	if err != nil {
		return nil
	}

	return &UserConsumer{
		ch:        ch,
		processor: processor,
	}
}

// TODO: решить вопрос с обработкой ошибок, паники
func (u UserConsumer) Start(ctx context.Context) {
	const op = "rabbitmq.Start"
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
				logger.Info("stopping event consumer by ctx")
				return
			case message := <-messages:
				logger.Info("Message: %s\n", message.Body)
				e := event(message.Body)
				err = u.processor.Process(ctx, *e)
				if err != nil || e == nil {
					logger.Warn(op, "err: %s\n", err)
					logger.Warn(op, "event: %s\n", e)
				}
			case <-sigchan:
				logger.Info("Interrupt detected!")
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

	logger.Info("event:", "e", e)

	return &e
}
