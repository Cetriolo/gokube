package main

import (
	"context"

	"log"

	"ride-sharing/services/trip-service/internal/domain"

	"net/http"

	h "ride-sharing/services/trip-service/internal/infrastructure/http"

	"ride-sharing/services/trip-service/internal/infrastructure/repository"

	"ride-sharing/services/trip-service/internal/service"

	"time"
)

func main() {
	ctx := context.Background()
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	mux := http.NewServeMux()
	fare := &domain.RideFareModel{
		UserID: "42",
	}
	httphandler := h.HttpHandler{Service: svc}
	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)

	}
	mux.HandleFunc("POST /preview", httphandler.HandleTripPreview)
	log.Println(t)
	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}
	// keep the program running for now
	for {
		time.Sleep(time.Second)
		if err := server.ListenAndServe(); err != nil {
			log.Printf("HTTP server error: %v", err)
		}
	}
}
