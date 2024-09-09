package app

import (
	"context"
	"log"

	"gitlab.com/acumen5524834/shop-service/internal/delivery/grpc/services"
	pb "gitlab.com/acumen5524834/shop-service/internal/pkg/genproto"
	"google.golang.org/protobuf/encoding/protojson"
)

type KafkaHandler struct {
	book    *services.BookRPC
	post    *services.PostRPC
	product *services.ProductRPC
}

func (h *KafkaHandler) BookCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.BookCreate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.book.CreateBook(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create book via Kafka: %v", err)
			return
		}
		log.Printf("Created book: %+v", res)
	}
}
func (h *KafkaHandler) UpdateCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.BookUpdate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.book.UpdateBook(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update book via Kafka: %v", err)
			return
		}
		log.Printf("Updated book: %+v", res)
	}
}
func (h *KafkaHandler) PostCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.PostCreate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.post.CreatePost(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create post via Kafka: %v", err)
			return
		}
		log.Printf("Created post: %+v", res)
	}
}
func (h *KafkaHandler) PostUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.PostUpdate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.post.UpdatePost(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update post via Kafka: %v", err)
			return
		}
		log.Printf("Updated post: %+v", res)
	}
}
func (h *KafkaHandler) ProductCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.ProductCreate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.product.CreateProduct(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create product via Kafka: %v", err)
			return
		}
		log.Printf("Created product: %+v", res)
	}
}
func (h *KafkaHandler) ProductUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.ProductUpdate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.product.UpdateProduct(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update product via Kafka: %v", err)
			return
		}
		log.Printf("Updated product: %+v", res)
	}
}
