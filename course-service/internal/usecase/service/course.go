package service

import (
	"context"

	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
	"gitlab.com/muallimah/course_service/internal/storage"
)

type CourseService struct {
	stg storage.StorageI
	pb.UnimplementedCourseServiceServer
}

func NewCourseService(stg storage.StorageI) *CourseService {
    return &CourseService{stg: stg}
}

func (s *CourseService) CreateCourse(c context.Context, req *pb.CourseCreateReq) (*pb.Void,error) {
    err :=  s.stg.Course().CreateCourse(req)
	return &pb.Void{}, err
}
func (s *CourseService) GetCourse(c context.Context, id *pb.ById) (*pb.Course, error) {
    return s.stg.Course().GetCourse(id)
}
func (s *CourseService) UpdateCourse(c context.Context, req *pb.CourseUpdateReq) (*pb.Void, error) {
    err := s.stg.Course().UpdateCourse(req)
	return &pb.Void{}, err
}
func (s *CourseService) DeleteCourse(c context.Context, id *pb.ById) (*pb.Void, error) {
    err := s.stg.Course().DeleteCourse(id)
    return &pb.Void{}, err
}
func (s *CourseService) ListCourses(c context.Context, req *pb.CourseListsReq) (*pb.CourseListsRes, error) {
    return s.stg.Course().ListCourses(req)
}