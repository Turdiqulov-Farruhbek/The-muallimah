package service

import (
	"context"

	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
	"gitlab.com/muallimah/course_service/internal/storage"
)


type UserLessonService struct {
	stg storage.StorageI
	pb.UnimplementedUserLessonServiceServer
}

func NewUserLessonService(stg storage.StorageI) *UserLessonService {
    return &UserLessonService{stg: stg}
}

func (s *UserLessonService) CreateUserLesson(c context.Context, req *pb.UserLessonCreateReq) (*pb.Void, error) {
    err := s.stg.UserLesson().CreateUserLesson(req)
    return &pb.Void{}, err
}
func (s *UserLessonService) GetUserLesson(c context.Context, id *pb.ById) (*pb.UserLesson, error) {
    return s.stg.UserLesson().GetUserLesson(id)
}
func (s *UserLessonService) UpdateUserLesson(c context.Context, req *pb.UserLessonUpdateReq) (*pb.Void, error) {
    err := s.stg.UserLesson().UpdateUserLesson(req)
    return &pb.Void{}, err
}
func (s *UserLessonService) DeleteUserLesson(c context.Context, id *pb.ById) (*pb.Void, error) {
    err := s.stg.UserLesson().DeleteUserLesson(id)
    return &pb.Void{}, err
}
func (s *UserLessonService) ListUserLessons(c context.Context, req *pb.UserLessonListsReq) (*pb.UserLessonListsRes, error) {
    return s.stg.UserLesson().ListUserLessons(req)
}