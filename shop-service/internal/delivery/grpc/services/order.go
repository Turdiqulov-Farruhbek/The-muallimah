package services

import (
	"context"

	"gitlab.com/acumen5524834/shop-service/internal/infrastructure/repository"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/genproto"
)

type OrderServiceClient struct {
	stg repository.Orders
	genproto.UnimplementedOrderServiceServer
}

func NewOrderServiceClient(stg repository.Orders) *OrderServiceClient {
	return &OrderServiceClient{stg: stg}
}

func (s *OrderServiceClient) CreateOrder(c context.Context, req *genproto.OrderCreateReq) (*genproto.Void, error) {
	_,err := s.stg.CreateOrder(c,req)
	return &genproto.Void{}, err
}
func (s *OrderServiceClient) GetOrder(c context.Context, id *genproto.OrderGetReq) (*genproto.OrderGetRes, error) {
	return s.stg.GetOrder(c,id)
}
func (s *OrderServiceClient) UpdateOrder(c context.Context, req *genproto.OrderUpdateReq) (*genproto.Void, error) {
	_,err := s.stg.UpdateOrder(c,req)
	return &genproto.Void{}, err
}
func (s *OrderServiceClient) DeleteOrder(c context.Context, id *genproto.OrderDeleteReq) (*genproto.Void, error) {
	_,err := s.stg.DeleteOrder(c,id)
	return &genproto.Void{}, err
}
func (s *OrderServiceClient) ListOrders(c context.Context, req *genproto.OrderListsReq) (*genproto.OrderListsRes, error) {
	return s.stg.ListOrders(c,req)
}
