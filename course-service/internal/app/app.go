package app

import (
	"errors"
	"log"
	"net"

	"gitlab.com/muallimah/course_service/internal/pkg/config"
	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
	"gitlab.com/muallimah/course_service/internal/pkg/postgres"
	"gitlab.com/muallimah/course_service/internal/storage/repo"
	"gitlab.com/muallimah/course_service/internal/usecase/kafka"
	"gitlab.com/muallimah/course_service/internal/usecase/minio"
	"gitlab.com/muallimah/course_service/internal/usecase/service"
	"google.golang.org/grpc"
)

func Run(cf *config.Config) {
	// connect to postgres
	pgm, err := postgres.New(cf)
	if err != nil {
		log.Fatal(err)
	}
	defer pgm.Close()

	// connect to kafka producer
	kf_p, err := kafka.NewKafkaProducer([]string{cf.KafkaUrl})
	if err != nil {
		log.Fatal(err)
	}
	// connect to minio
	minio, err := minio.MinIOConnect(cf)
	if err != nil {
		log.Fatal(err)
	}

	// repo
	db := repo.NewStorage(pgm.DB)

	// register kafka handlers
	k_handler := KafkaHandler{
		certificate: service.NewCertificateService(db, kf_p, minio),
	}

	if err := Registries(&k_handler, cf); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", cf.GRPCPort)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	// set grpc server
	server := grpc.NewServer()
	pb.RegisterCategoryServiceServer(server, service.NewCategoryService(db))
	pb.RegisterCertificateServiceServer(server, service.NewCertificateService(db, kf_p, minio))
	pb.RegisterCourseServiceServer(server, service.NewCourseService(db))
	pb.RegisterLessonServiceServer(server, service.NewLessonService(db))
	pb.RegisterMaterialServiceServer(server, service.NewMaterialService(db))
	pb.RegisterUserCourseServiceServer(server, service.NewUserCourseService(db))
	pb.RegisterUserLessonServiceServer(server, service.NewUserLessonService(db))
	pb.RegisterTransactionServiceServer(server, service.NewTransactionService(db))

	// start server

	log.Println("Server started on", cf.GRPCPort)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	defer lis.Close()
}

func Registries(k_handler *KafkaHandler, cfg *config.Config) error {
	brokers := []string{cfg.KafkaUrl}
	kcm := kafka.NewKafkaConsumerManager()

	if err := kcm.RegisterConsumer(brokers, "certificate-update", "certificate-u", k_handler.CertificateUpdate()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'certificate-update' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	return nil
}
