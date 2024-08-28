package service

import (
	"context"
	"errors"
	"fmt"

	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
	"gitlab.com/muallimah/course_service/internal/storage"
	"gitlab.com/muallimah/course_service/internal/usecase/kafka"
	"gitlab.com/muallimah/course_service/internal/usecase/minio"
	pdfmaker "gitlab.com/muallimah/course_service/internal/usecase/pdf_maker"
)

type CertificateService struct {
	stg      storage.StorageI
	producer kafka.KafkaProducer
	minio    *minio.MinIO
	pb.UnimplementedCertificateServiceServer
}

func NewCertificateService(stg storage.StorageI) *CertificateService {
	return &CertificateService{stg: stg}
}

func (s *CertificateService) CreateCertificate(c context.Context, id *pb.ById) (*pb.Void, error) {
	return s.stg.Certificate().CreateCertificate(id)
}
func (s *CertificateService) GetCertificate(c context.Context, id *pb.ById) (*pb.CertificateGet, error) {
	return s.stg.Certificate().GetCertificate(id)
}
func (s *CertificateService) UpdateCertificate(c context.Context, req *pb.CertificateUpdate) (*pb.Void, error) {
	///  DO NOT FORGET TO ADD API TO CREATE CERTIFICATE!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	certificate, err := s.stg.Certificate().GetCertificate(&pb.ById{Id: req.Id})
	if err != nil {
		return nil, errors.New("internal usecases certificate get certificate error: " + err.Error())
	}
	if req.Body.Status == "confirmed" {
		res, err := pdfmaker.GenerateCertificate(
			certificate.UserCourse.User.FirstName,
			certificate.UserCourse.Course.Name,
			"The Muallimal",
			s.minio.Cf,
		)
		if err != nil {
			return nil, errors.New("internal - usecases- service -certificcate -genrating certificate error: " + err.Error())
		}
		cer_url, err := s.minio.Upload(*res, ".pdf")
		if err != nil {
			return nil, errors.New("internal usecases minio upload error: " + err.Error())
		}
		req.Body.FileUrl = *cer_url
	}
	// write to kafka
	message := fmt.Sprintf("dear %s your certificate request has been %s",certificate.UserCourse.User.FirstName,req.Body.Status)
	input,err := 

	return s.stg.Certificate().UpdateCertificate(req)
}
func (s *CertificateService) DeleteCertificate(c context.Context, id *pb.ById) (*pb.Void, error) {
	return s.stg.Certificate().DeleteCertificate(id)
}
func (s *CertificateService) ListCertificates(c context.Context, req *pb.CertificateFilter) (*pb.CertificateList, error) {
	return s.stg.Certificate().ListCertificates(req)
}
