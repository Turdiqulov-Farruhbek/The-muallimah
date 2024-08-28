package app

import (
	"log"
	"net"

	"gitlab.com/muallimah/course_service/internal/pkg/config"
	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
	"gitlab.com/muallimah/course_service/internal/pkg/postgres"
	"gitlab.com/muallimah/course_service/internal/storage/repo"
	// "gitlab.com/muallimah/course_service/internal/usecase/kafka"
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
	// kf_p, err := kafka.NewKafkaProducer([]string{cf.KafkaUrl})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// repo
	db := repo.NewStorage(pgm.DB)

	// ch := make(chan int, 3)
	// set grpc server
	server := grpc.NewServer()

	// register service
	pb.RegisterCategoryServiceServer(server, service.NewCategoryService(db))
	pb.RegisterCertificateServiceServer(server, service.NewCertificateService(db))
	pb.RegisterCourseServiceServer(server, service.NewCourseService(db))
	pb.RegisterLessonServiceServer(server, service.NewLessonService(db))
	pb.RegisterMaterialServiceServer(server, service.NewMaterialService(db))
	pb.RegisterUserCourseServiceServer(server, service.NewUserCourseService(db))
	pb.RegisterUserLessonServiceServer(server, service.NewUserLessonService(db))

	// start server
	lis, err := net.Listen("tcp", cf.GRPCPort)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	defer lis.Close()

	log.Println("Server started on", cf.GRPCPort)
	select {} // wait forever until killed
}
