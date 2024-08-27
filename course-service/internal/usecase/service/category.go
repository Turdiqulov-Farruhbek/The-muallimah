package service

import (
	"context"

	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
	"gitlab.com/muallimah/course_service/internal/storage"
)


type CategoryService struct {
	stg storage.StorageI
	pb.UnimplementedCategoryServiceServer
}

func NewCategoryService(stg storage.StorageI) *CategoryService {
    return &CategoryService{stg: stg}
}
func (s *CategoryService) CreateCategory(c context.Context, req *pb.CategoryCreateReq) (*pb.Void, error) {
	return s.stg.Category().CreateCategory(req)
}
func (s *CategoryService) GetCategory(c context.Context, req *pb.ById) (*pb.Category, error) {
    return s.stg.Category().GetCategory(req)
}
func (s *CategoryService) ListCategories(c context.Context, req *pb.Pagination) (*pb.CategoryListRes, error) {
    return s.stg.Category().ListCategories(req)
}
func (s *CategoryService) UpdateCategory(c context.Context, req *pb.CategoryUpdateReq) (*pb.Void, error) {
    return s.stg.Category().UpdateCategory(req)
}
func (s *CategoryService) DeleteCategory(c context.Context, req *pb.ById) (*pb.Void, error) {
    return s.stg.Category().DeleteCategory(req)
}