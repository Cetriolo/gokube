package events

import (
	"context"
	"encoding/json"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/messaging"
)

type TripEventPublisher struct {
	rabbit, q *messaging.RabbitMQ
}

func NewTripEventPublisher(rabbit *messaging.RabbitMQ) *TripEventPublisher {
	return &TripEventPublisher{rabbit: rabbit}
}

func (p *TripEventPublisher) PublishTripCreated(ctx context.Context, trip *domain.TripModel) error {
	payload := messaging.TripEventData{
		Trip: trip.ToProto(),
	}
	tripEventJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return p.rabbit.PublishMessage(ctx, contracts.TripEventCreated, contracts.AmqpMessage{
		OwnerID: trip.UserID,
		Data:    tripEventJson,
	})

}
