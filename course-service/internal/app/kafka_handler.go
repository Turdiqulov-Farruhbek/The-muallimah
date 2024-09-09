package app

import (
	"context"
	"log"

	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
	"gitlab.com/muallimah/course_service/internal/usecase/service"
	"google.golang.org/protobuf/encoding/protojson"
)

type KafkaHandler struct {
	certificate *service.CertificateService
}

func (h *KafkaHandler) CertificateUpdate() func(message []byte) {
	return func(message []byte) {

		//unmarshal the message
		var cer pb.CertificateUpdate
		if err := protojson.Unmarshal(message, &cer); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := h.certificate.UpdateCertificate(context.Background(), &cer)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Updated Certificate: %+v", res)
	}
}
