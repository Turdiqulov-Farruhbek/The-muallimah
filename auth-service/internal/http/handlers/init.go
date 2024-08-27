package handlers

import (
	"context"

	pb "auth-service/internal/pkg/genproto"
	minio_connect "auth-service/internal/pkg/minio"
	rdb "auth-service/internal/pkg/redis"

	"google.golang.org/grpc"
)

type HTTPHandler struct {
	RDB   *rdb.RedisClient
	Minio *minio_connect.MinioClient
	US    pb.UserServiceClient
}

func NewHandler(us *grpc.ClientConn) *HTTPHandler {
	db, err := rdb.NewRedisClient(context.Background())
	if err != nil {
		panic(err)
	}

	mc, err := minio_connect.NewMinioClient()
	if err != nil {
		panic(err)
	}

	return &HTTPHandler{
		RDB:   db,
		US:    pb.NewUserServiceClient(us),
		Minio: mc,
	}
}
