package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"ride-sharing/shared/env"
	"ride-sharing/shared/messaging"
	"syscall"

	grpcserver "google.golang.org/grpc"
)

var GrpcAddr = ":9092"

func main() {
	rabbitMQString := env.GetString("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672/")

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

	grpcServer := grpcserver.NewServer()

	log.Println("Starting gRPC server Driver service on port %s " + lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			cancel()
		}
	}()
	<-ctx.Done()
	log.Println("Shutting down gRPC server Driver service")
	grpcServer.GracefulStop()
	log.Println("gRPC server Driver service stopped")
}
