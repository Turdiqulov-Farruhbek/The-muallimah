package services

import (
	"context"

	"gitlab.com/acumen5524834/shop-service/internal/infrastructure/repository"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/genproto"
)

type OrderService struct {
	stg repository.Orders
	genproto.UnimplementedOrderServiceServer
}

func NewOrderService(stg repository.Orders) *OrderService {
	return &OrderService{stg: stg}
}

func (s *OrderService) CreateOrder(c context.Context, req *genproto.OrderCreateReq) (*genproto.Void, error) {
	_,err := s.stg.CreateOrder(c,req)
	return &genproto.Void{}, err
}
func (s *OrderService) GetOrder(c context.Context, id *genproto.OrderGetReq) (*genproto.OrderGetRes, error) {
	return s.stg.GetOrder(c,id)
}
func (s *OrderService) UpdateOrder(c context.Context, req *genproto.OrderUpdateReq) (*genproto.Void, error) {
	_,err := s.stg.UpdateOrder(c,req)
	return &genproto.Void{}, err
}
func (s *OrderService) DeleteOrder(c context.Context, id *genproto.OrderDeleteReq) (*genproto.Void, error) {
	_,err := s.stg.DeleteOrder(c,id)
	return &genproto.Void{}, err
}
func (s *OrderService) ListOrders(c context.Context, req *genproto.OrderListsReq) (*genproto.OrderListsRes, error) {
	return s.stg.ListOrders(c,req)
}
