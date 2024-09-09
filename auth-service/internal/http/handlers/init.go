package handlers

import (
	"context"

	pb "auth-service/internal/pkg/genproto"
	minio_connect "auth-service/internal/pkg/minio"
	rdb "auth-service/internal/pkg/redis"

	l "github.com/azizbek-qodirov/logger"
	"google.golang.org/grpc"
)

type HTTPHandler struct {
	Logger *l.Logger
	RDB    *rdb.RedisClient
	Minio  *minio_connect.MinioClient
	US     pb.UserServiceClient
}

func NewHandler(us *grpc.ClientConn, logger *l.Logger) *HTTPHandler {
	db, err := rdb.NewRedisClient(context.Background())
	if err != nil {
		logger.ERROR.Panicln("Redis not connected due to error: " + err.Error())
	}

	mc, err := minio_connect.NewMinioClient()
	if err != nil {
		logger.ERROR.Panicln("Minio not connected due to error: " + err.Error())
	}

	return &HTTPHandler{
		Logger: logger,
		RDB:    db,
		US:     pb.NewUserServiceClient(us),
		Minio:  mc,
	}
}
