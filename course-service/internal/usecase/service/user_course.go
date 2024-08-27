package service

import (
	"context"

	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
	"gitlab.com/muallimah/course_service/internal/storage"
)

type UserCourseService struct {
	stg storage.StorageI
	pb.UnimplementedUserCourseServiceServer
}

func NewUserCourseService(stg storage.StorageI) *UserCourseService {
    return &UserCourseService{stg: stg}
}

func (s *UserCourseService) EnrollUserInCourse(c context.Context, req *pb.UserCourseCreateReq) (*pb.Void, error) {
    err := s.stg.UserCourse().EnrollUserInCourse(req)
	return &pb.Void{}, err
}
func (s *UserCourseService) GetUserCourse(c context.Context, id *pb.ById) (*pb.UserCourse, error) {
    return s.stg.UserCourse().GetUserCourse(id)
}
func (s *UserCourseService) UpdateUserCourse(c context.Context, req *pb.UserCourseUpdateReq) (*pb.Void, error) {
    err := s.stg.UserCourse().UpdateUserCourse(req)
    return &pb.Void{}, err
}
func (s *UserCourseService) DeleteUserCourse(c context.Context, id *pb.ById) (*pb.Void, error) {
    err := s.stg.UserCourse().DeleteUserCourse(id)
    return &pb.Void{}, err
}
func (s *UserCourseService) ListUserCourses(c context.Context, req *pb.UserCourseListsReq) (*pb.UserCourseListsRes, error) {
    return s.stg.UserCourse().ListUserCourses(req)
}