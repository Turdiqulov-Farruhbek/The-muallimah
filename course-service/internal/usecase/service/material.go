package service

import (
	"context"

	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
	"gitlab.com/muallimah/course_service/internal/storage"
)

type MaterialService struct {
	stg storage.StorageI
	pb.UnimplementedMaterialServiceServer
}

func NewMaterialService(stg storage.StorageI) *MaterialService {
    return &MaterialService{stg: stg}
}

func (s *MaterialService) CreateMaterial(c context.Context, req *pb.MaterialCreateReq) (*pb.Void, error) {
	return s.stg.Material().CreateMaterial(req)
}
func (s *MaterialService) GetMaterial(c context.Context, id *pb.ById) (*pb.MaterialGetRes, error) {
    return s.stg.Material().GetMaterial(id)
}
func (s *MaterialService) UpdateMaterial(c context.Context, req *pb.MaterialUpdateReq) (*pb.Void, error) {
    return s.stg.Material().UpdateMaterial(req)
}
func (s *MaterialService) DeleteMaterial(c context.Context, id *pb.ById) (*pb.Void, error) {
    return s.stg.Material().DeleteMaterial(id)
}
func (s *MaterialService) ListMaterials(c context.Context, req *pb.MaterialListReq) (*pb.MaterialListRes, error) {
    return s.stg.Material().ListMaterials(req)
}