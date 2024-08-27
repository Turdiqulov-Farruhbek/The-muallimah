package grpc_service_clients

import (
	"fmt"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/genproto"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/config"

	"google.golang.org/grpc"
)

type ServiceClients interface {
	BookService() genproto.BookServiceClient
	PostService() genproto.PostServiceClient
	ProductService() genproto.ProductServiceClient
	Close()
}

type serviceClients struct {
	bookService genproto.BookServiceClient
	postService genproto.PostServiceClient
	productService genproto.ProductServiceClient
	services    []*grpc.ClientConn
}

func New(config *config.Config) (ServiceClients, error) {
	shopServiceConnection, err := grpc.Dial(
		fmt.Sprintf("%s%s", config.DB.Host, config.DB.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &serviceClients{
		bookService: genproto.NewBookServiceClient(shopServiceConnection),
		postService: genproto.NewPostServiceClient(shopServiceConnection),
		productService: genproto.NewProductServiceClient(shopServiceConnection),
		services: []*grpc.ClientConn{
			shopServiceConnection,
		},
	}, nil
}

func (s *serviceClients) Close() {
	// closing store-management service
	for _, conn := range s.services {
		err := conn.Close()
		if err != nil {
			return
		}
	}
}

func (s *serviceClients) BookService() genproto.BookServiceClient {
	return s.bookService
}
func (s *serviceClients) PostService() genproto.PostServiceClient {
	return s.postService
}
func (s *serviceClients) ProductService() genproto.ProductServiceClient {
	return s.productService
}

