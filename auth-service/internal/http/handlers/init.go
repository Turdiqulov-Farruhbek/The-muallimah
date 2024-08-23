package handlers

import (
	pb "auth-service/internal/pkg/genproto"

	"google.golang.org/grpc"
)

type HTTPHandler struct {
	US pb.UserServiceClient
}

func NewHandler(us *grpc.ClientConn) *HTTPHandler {
	return &HTTPHandler{US: pb.NewUserServiceClient(us)}
}
