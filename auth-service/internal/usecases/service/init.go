package service

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/storage/repo"
	"fmt"
	"log"
	"net"

	pb "auth-service/internal/pkg/genproto"

	"google.golang.org/grpc"
)

func Server(cfg *config.Config, stg *repo.Storage) {
	userService := NewUserService(stg)

	newServer := grpc.NewServer()
	pb.RegisterUserServiceServer(newServer, userService)

	lis, err := net.Listen("tcp", cfg.AUTH_GRPC_PORT)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("gRPC server is running on port ", cfg.AUTH_GRPC_PORT)
	err = newServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
