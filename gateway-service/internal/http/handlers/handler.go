package handlers

import (
	"muallimah-gateway/internal/grpc"
	"muallimah-gateway/internal/pkg/kafka"
	"muallimah-gateway/internal/pkg/logger"

	// "github.com/go-redis/redis"
)

type Handler struct {
	Clients  grpc.Clients
	Producer kafka.KafkaProducer
	Logger   *logger.Logger
}

func NewHandler(clients grpc.Clients, producer kafka.KafkaProducer, logger *logger.Logger) *Handler {
	return &Handler{Clients: clients, Producer: producer, Logger: logger}
}
