package app

import (
	"log"
	"path/filepath"
	"runtime"

	"muallimah-gateway/internal/grpc"
	"muallimah-gateway/internal/http"
	"muallimah-gateway/internal/http/handlers"
	"muallimah-gateway/internal/pkg/config"
	"muallimah-gateway/internal/pkg/kafka"
	"muallimah-gateway/internal/pkg/logger"

	// "github.com/go-redis/redis"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func Run(cfg config.Config) {
	logger := logger.NewLogger(basepath, cfg.LogPath)
	clients, err := grpc.NewClients(&cfg)
	if err != nil {
		logger.ERROR.Println("Failed to create gRPC clients", err)
		log.Fatal(err)
		return
	}

	//connect to kafka
	broker := []string{cfg.KafkaUrl}
	kafka, err := kafka.NewKafkaProducer(broker)
	if err != nil {
		logger.ERROR.Println("Failed to connect to Kafka", err)
		log.Fatal(err)
		return
	}
	defer kafka.Close()


	// make handler
	h := handlers.NewHandler(*clients, kafka, logger)

	// make gin
	router := http.NewGin(h)

	// start server
	router.Run(":5050")
}
