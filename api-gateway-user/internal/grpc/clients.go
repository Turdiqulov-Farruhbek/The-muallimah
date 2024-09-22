package grpc

import (
	"muallimah-gateway/internal/pkg/config"
	pb "muallimah-gateway/internal/pkg/genproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Material     pb.MaterialServiceClient
	Book         pb.BookServiceClient
	Category     pb.CategoryServiceClient
	Certificate  pb.CertificateServiceClient
	Course       pb.CourseServiceClient
	Evaluation   pb.EvaluationServiceClient
	Lesson       pb.LessonServiceClient
	Order        pb.OrderServiceClient
	Post         pb.PostServiceClient
	Product      pb.ProductServiceClient
	Transaction  pb.TransactionServiceClient
	User         pb.UserServiceClient
	UserLesson   pb.UserLessonServiceClient
	UserCourse   pb.UserCourseServiceClient
	Notification pb.NotificationServiceClient
}

func NewClients(cfg *config.Config) (*Clients, error) {
	auth_conn, err := grpc.NewClient(cfg.AuthUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	course_conn, err := grpc.NewClient(cfg.CoursetUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	shop_conn, err := grpc.NewClient(cfg.ShopUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	notif_conn, err := grpc.NewClient(cfg.NotificationUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	userClient := pb.NewUserServiceClient(auth_conn)

	bookClient := pb.NewBookServiceClient(shop_conn)
	productClient := pb.NewProductServiceClient(shop_conn)
	postClient := pb.NewPostServiceClient(shop_conn)

	courseClient := pb.NewCourseServiceClient(course_conn)
	materialClient := pb.NewMaterialServiceClient(course_conn)
	categoryClient := pb.NewCategoryServiceClient(course_conn)
	certificateClient := pb.NewCertificateServiceClient(course_conn)
	userCourseClient := pb.NewUserCourseServiceClient(course_conn)
	evaluationClient := pb.NewEvaluationServiceClient(course_conn)
	lessonClient := pb.NewLessonServiceClient(course_conn)
	orderClient := pb.NewOrderServiceClient(course_conn)
	transactionClient := pb.NewTransactionServiceClient(course_conn)
	userLessonClient := pb.NewUserLessonServiceClient(course_conn)

	notificationClient := pb.NewNotificationServiceClient(notif_conn)

	return &Clients{
		Material:     materialClient,
		Book:         bookClient,
		Category:     categoryClient,
		Certificate:  certificateClient,
		Course:       courseClient,
		Evaluation:   evaluationClient,
		Lesson:       lessonClient,
		Order:        orderClient,
		Post:         postClient,
		Product:      productClient,
		Transaction:  transactionClient,
		User:         userClient,
		UserLesson:   userLessonClient,
		UserCourse:   userCourseClient,
		Notification: notificationClient,
	}, nil
}
