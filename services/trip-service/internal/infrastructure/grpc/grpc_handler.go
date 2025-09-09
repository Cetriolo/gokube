package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"log"
	"ride-sharing/services/trip-service/internal/domain"
	pb "ride-sharing/shared/proto/trip"
	"ride-sharing/shared/types"
)

type gRPCHandler struct {
	pb.UnimplementedTripServiceServer
	service domain.TripService
}

func NewgRPCHandler(server *grpc.Server, service domain.TripService) *gRPCHandler {

	handler := &gRPCHandler{
		service: service,
	}
	pb.RegisterTripServiceServer(server, handler)
	return handler
}

func (h *gRPCHandler) CreateTrip(ctx context.Context, req *pb.CreateTripRequest) (*pb.CreateTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrip not implemented")
}

func (h *gRPCHandler) PreviewTrip(ctx context.Context, req *pb.PreviewTripRequest) (*pb.PreviewTripResponse, error) {

	pickup := req.GetStartLocation()
	destination := req.GetEndLocation()

	pickUpCoord := &types.Coordinate{Latitude: pickup.GetLatitude(),
		Longitude: pickup.GetLongitude()}
	destCoord := &types.Coordinate{Latitude: destination.GetLatitude(),
		Longitude: destination.GetLongitude()}

	t, err := h.service.GetRoute(ctx, pickUpCoord, destCoord)
	if err != nil {
		log.Printf("GetRoute err: %v", err)
		return nil, status.Errorf(codes.Internal, "GetRoute err: %v", err)
	}

	return &pb.PreviewTripResponse{
		Route:     t.ToProto(),
		RideFares: []*pb.RideFare{},
	}, nil
}
