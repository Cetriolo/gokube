package messaging

import pb "ride-sharing/shared/proto/trip"

const (
	FindAvaiableDriversQueue = "find_available_drivers"
)

type TripEventData struct {
	Trip *pb.Trip `json:"trip"`
}
