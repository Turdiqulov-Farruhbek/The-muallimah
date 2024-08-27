package service

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/storage/repo"
	"net"

	pb "auth-service/internal/pkg/genproto"

	l "github.com/azizbek-qodirov/logger"

	"google.golang.org/grpc"
)

func Server(cfg *config.Config, logger l.Logger, stg *repo.Storage) {
	userService := NewUserService(stg)

	newServer := grpc.NewServer()
	pb.RegisterUserServiceServer(newServer, userService)

	lis, err := net.Listen("tcp", cfg.AUTH_GRPC_PORT)
	if err != nil {
		logger.ERROR.Panicln("Failed to listen gRPC server: ", err.Error())
	}

	logger.INFO.Println("gRPC server is running on port ", cfg.AUTH_GRPC_PORT)
	err = newServer.Serve(lis)
	if err != nil {
		logger.ERROR.Panicln("Failed to serve gRPC server: ", err.Error())
	}
}
