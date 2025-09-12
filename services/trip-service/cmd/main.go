package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"ride-sharing/services/trip-service/internal/infrastructure/events"
	"ride-sharing/services/trip-service/internal/infrastructure/grpc"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/env"
	"syscall"

	grpcserver "google.golang.org/grpc"
	"ride-sharing/shared/messaging"
)

var GrpcAddr = ":9093"

func main() {
	rabbitMQString := env.GetString("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672/")
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	lis, err := net.Listen("tcp", GrpcAddr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//RabbitMQ connection
	rabbitmq, err := messaging.NewRabbitMQ(rabbitMQString)
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitmq.Close()
	log.Println("Staring rabbitmq connection")
	// Create gRPC server

	publisher := events.NewTripEventPublisher(rabbitmq)

	grpcServer := grpcserver.NewServer()

	grpc.NewgRPCHandler(grpcServer, svc, publisher)

	log.Println("Starting gRPC server Trip service on port %s " + lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			cancel()
		}
	}()
	<-ctx.Done()
	log.Println("Shutting down gRPC server Trip service")
	grpcServer.GracefulStop()
	log.Println("gRPC server Trip service stopped")
}
