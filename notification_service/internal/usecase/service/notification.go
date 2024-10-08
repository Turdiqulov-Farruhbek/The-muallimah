package service

import (
	"context"
	"errors"

	pb "gitlab.com/muallimah/notification_service/internal/pkg/genproto"
	"gitlab.com/muallimah/notification_service/internal/storage"
)

type NotificationService struct {
	stg storage.StorageI
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationService(stg storage.StorageI) *NotificationService {
	return &NotificationService{stg: stg}
}

func (s *NotificationService) CreateNotification(ctx context.Context, req *pb.NotificationCreate) (*pb.Void, error) {
	return s.stg.Notification().CreateNotification(req)
}
func (s *NotificationService) NotifyAll(ctx context.Context, req *pb.NotificationMessage) (*pb.Void, error) {
	return s.stg.Notification().NotifyAll(req)
}
func (s *NotificationService) DeleteNotificcation(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	if req.Id == "" || req.Id == "string" || req.Id == " " {
		return nil, errors.New("invalid request, id is required")
	}
	return s.stg.Notification().DeleteNotificcation(req)
}
func (s *NotificationService) UpdateNotificcation(ctx context.Context, req *pb.NotificationUpdate) (*pb.Void, error) {
	if req.NotificationId == "" || req.NotificationId == "string" || req.NotificationId == " " {
		return nil, errors.New("invalid request, notification_id is required")
	}
	return s.stg.Notification().UpdateNotificcation(req)
}
func (s *NotificationService) GetNotifications(ctx context.Context, req *pb.NotifFilter) (*pb.NotificationList, error) {
	return s.stg.Notification().GetNotifications(req)
}
func (s *NotificationService) GetNotification(ctx context.Context, req *pb.ById) (*pb.NotificationGet, error) {
	return s.stg.Notification().GetNotification(req)
}
