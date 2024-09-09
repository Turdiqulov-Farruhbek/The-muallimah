package usecase

import (
	"context"
	"time"

	"gitlab.com/acumen5524834/shop-service/internal/infrastructure/repository"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/genproto"
)

type Orders interface {
	CreateOrder(ctx context.Context, order *genproto.OrderCreateReq) (*genproto.Void, error)
	UpdateOrder(ctx context.Context, order *genproto.OrderUpdateReq) (*genproto.Void, error)
	DeleteOrder(ctx context.Context, orderID *genproto.OrderDeleteReq) (*genproto.Void, error)
	ListOrders(ctx context.Context, order *genproto.OrderListsReq) (*genproto.OrderListsRes, error)
	GetOrder(ctx context.Context, orderID *genproto.OrderGetReq) (*genproto.OrderGetRes, error)
}
type orderService struct {
	repo repository.Orders
}

func NewOrderService(repo repository.Orders) Orders {
	return &orderService{
		repo: repo,
	}
}

func (u orderService) CreateOrder(ctx context.Context, order *genproto.OrderCreateReq) (*genproto.Void, error) {
	ctx, cancel := context.WithTimeout(ctx,time.Duration(10)*time.Second)
	defer cancel()

	return u.repo.CreateOrder(ctx, order)
}

func (u orderService) UpdateOrder(ctx context.Context, order *genproto.OrderUpdateReq) (*genproto.Void, error) {
	ctx, cancel := context.WithTimeout(ctx,time.Duration(10)*time.Second)
	defer cancel()

	return u.repo.UpdateOrder(ctx, order)
}

func (u orderService) DeleteOrder(ctx context.Context, orderID *genproto.OrderDeleteReq) (*genproto.Void, error) {
	ctx, cancel := context.WithTimeout(ctx,time.Duration(10)*time.Second)
	defer cancel()

	return u.repo.DeleteOrder(ctx, orderID)
}

func (u orderService) ListOrders(ctx context.Context, order *genproto.OrderListsReq) (*genproto.OrderListsRes, error) {
	ctx, cancel := context.WithTimeout(ctx,time.Duration(10)*time.Second)
	defer cancel()

	return u.repo.ListOrders(ctx, order)
}

func (u orderService) GetOrder(ctx context.Context, orderID *genproto.OrderGetReq) (*genproto.OrderGetRes, error) {
	ctx, cancel := context.WithTimeout(ctx,time.Duration(10)*time.Second)
	defer cancel()

	return u.repo.GetOrder(ctx, orderID)
}
