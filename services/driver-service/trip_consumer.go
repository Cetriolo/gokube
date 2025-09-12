package main

import (
	"context"
	"log"
	"ride-sharing/shared/messaging"

	amqp "github.com/rabbitmq/amqp091-go"
)

type tripConsumer struct {
	rabbitmq *messaging.RabbitMQ
}

func NewTripConsumer(rabbitmq *messaging.RabbitMQ) *tripConsumer {
	return &tripConsumer{rabbitmq: rabbitmq}
}

func (c *tripConsumer) Listen() error {
	return c.rabbitmq.ConsumeMessages("hello", func(ctx context.Context, d amqp.Delivery) error {
		log.Printf("Received a message: %s", d.Body)
		return nil
	})
}
