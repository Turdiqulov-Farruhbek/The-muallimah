package app

import (
	"fmt"
	"log"
	"net"

	"gitlab.com/muallimah/notification_service/internal/pkg/config"
	"gitlab.com/muallimah/notification_service/internal/pkg/postgres"

	pb "gitlab.com/muallimah/notification_service/internal/pkg/genproto"
	"gitlab.com/muallimah/notification_service/internal/pkg/kafka"
	repo "gitlab.com/muallimah/notification_service/internal/storage/repo"
	"gitlab.com/muallimah/notification_service/internal/usecase/service"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config) {
	//connect to db
	db_m,err := postgres.New(cfg)
	if err!= nil {
        log.Fatal(err)
    }
	defer db_m.Close()
	///
	stg,err := repo.NewStorage(db_m.DB,cfg)

	//
	notificationService := service.NewNotificationService(stg)

	//kafka\\*//\\\
	brokers := []string{cfg.KafkaHost + cfg.KafkaPort}

	kcm := kafka.NewKafkaConsumerManager()

	if err := kcm.RegisterConsumer(brokers, "notification-create", "notification", kafka.NotificationCreateHandler(notificationService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'notification-create' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}
	if err := kcm.RegisterConsumer(brokers, "notify-all", "notification-all", kafka.NotifyAllHandler(notificationService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'notify-all' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}
	server := grpc.NewServer()
	pb.RegisterNotificationServiceServer(server, notificationService)

	lis, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running on " + cfg.HTTPPort)
	err = server.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
