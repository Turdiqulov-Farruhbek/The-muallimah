package service

import (
	"context"

	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
	"gitlab.com/muallimah/course_service/internal/storage"
)

type CertificateService struct {
	stg storage.StorageI
	pb.UnimplementedCertificateServiceServer
}

func(s *CertificateService) CreateCertificate(c context.Context, id *pb.ById) (*pb.Void,error) {
    return s.stg.Certificate().CreateCertificate(id)
}
func (s *CertificateService) GetCertificate(c context.Context, id *pb.ById) (*pb.CertificateGet, error) {
    return s.stg.Certificate().GetCertificate(id)
}
func (s *CertificateService) UpdateCertificate(c context.Context, req *pb.CertificateUpdate) (*pb.Void, error) {
	///  DO NOT FORGET TO ADD API TO CREATE CERTIFICATE!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
    return s.stg.Certificate().UpdateCertificate(req)
}
func (s *CertificateService) DeleteCertificate(c context.Context, id *pb.ById) (*pb.Void, error) {
    return s.stg.Certificate().DeleteCertificate(id)
}
func (s *CertificateService) ListCertificates(c context.Context, req *pb.CertificateFilter) (*pb.CertificateList, error) {
    return s.stg.Certificate().ListCertificates(req)
}