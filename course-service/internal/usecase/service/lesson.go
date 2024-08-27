package service

import (
	"context"

	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
	"gitlab.com/muallimah/course_service/internal/storage"
)

type LessonService struct {
	stg storage.StorageI
	pb.UnimplementedLessonServiceServer
}

func NewLessonService(stg storage.StorageI) *LessonService {
    return &LessonService{stg: stg}
}

func (s *LessonService) CreateLesson(c context.Context, req *pb.LessonCreateReq) (*pb.Void, error) {
    err := s.stg.Lesson().CreateLesson(req)
    return &pb.Void{}, err
}
func (s *LessonService) GetLesson(c context.Context, id *pb.ById) (*pb.LessonGet, error) {
    return s.stg.Lesson().GetLesson(id)
}
func (s *LessonService) UpdateLesson(c context.Context, req *pb.LessonUpdate) (*pb.Void, error) {
    err := s.stg.Lesson().UpdateLesson(req)
    return &pb.Void{}, err
}
func (s *LessonService) DeleteLesson(c context.Context, id *pb.ById) (*pb.Void, error) {
    err := s.stg.Lesson().DeleteLesson(id)
    return &pb.Void{}, err
}
func (s *LessonService) ListLessons(c context.Context, req *pb.LessonFilter) (*pb.LessonList, error) {
    return s.stg.Lesson().ListLessons(req)
}