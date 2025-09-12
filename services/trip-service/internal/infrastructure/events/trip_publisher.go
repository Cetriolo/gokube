package events

import (
	"context"
	"ride-sharing/shared/messaging"
)

type TripEventPublisher struct {
	rabbit, q *messaging.RabbitMQ
}

func NewTripEventPublisher(rabbit *messaging.RabbitMQ) *TripEventPublisher {
	return &TripEventPublisher{rabbit: rabbit}
}

func (p *TripEventPublisher) PublishTripCreated(ctx context.Context) error {
	return p.rabbit.PublishMessage(ctx, "hello", "Trip Created")

}
