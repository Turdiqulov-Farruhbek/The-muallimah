package app

import (
	"errors"

	"gitlab.com/acumen5524834/shop-service/internal/pkg/config"
	"gitlab.com/acumen5524834/shop-service/internal/usecase/kafka"
)

func Register(h *KafkaHandler, cfg *config.Config) error {

	brokers := []string{cfg.KafkaUrl}
	kcm := kafka.NewKafkaConsumerManager()

	if err := kcm.RegisterConsumer(brokers, "book-create", "book-create-id", h.BookCreateHandler()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'book-create' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "book-update", "book-update-id", h.UpdateCreateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'book-update' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "post-create", "post-create-id", h.PostCreateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'post-create' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "post-update", "post-update-id", h.PostUpdateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'post-update' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "product-create", "product-create-id", h.ProductCreateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'product-create' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "product-update", "product-update-id", h.ProductUpdateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'product-update' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}

	return nil
}