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
type OrderServiceClient struct {
	repo repository.Orders
	genproto.UnimplementedOrderServiceServer
}

func NewOrderService(repo repository.Orders) Orders {
	return &OrderServiceClient{
		repo: repo,
	}
}

func (u OrderServiceClient) CreateOrder(ctx context.Context, order *genproto.OrderCreateReq) (*genproto.Void, error) {
	ctx, cancel := context.WithTimeout(ctx,time.Duration(10)*time.Second)
	defer cancel()

	return u.repo.CreateOrder(ctx, order)
}

func (u OrderServiceClient) UpdateOrder(ctx context.Context, order *genproto.OrderUpdateReq) (*genproto.Void, error) {
	ctx, cancel := context.WithTimeout(ctx,time.Duration(10)*time.Second)
	defer cancel()

	return u.repo.UpdateOrder(ctx, order)
}

func (u OrderServiceClient) DeleteOrder(ctx context.Context, orderID *genproto.OrderDeleteReq) (*genproto.Void, error) {
	ctx, cancel := context.WithTimeout(ctx,time.Duration(10)*time.Second)
	defer cancel()

	return u.repo.DeleteOrder(ctx, orderID)
}

func (u OrderServiceClient) ListOrders(ctx context.Context, order *genproto.OrderListsReq) (*genproto.OrderListsRes, error) {
	ctx, cancel := context.WithTimeout(ctx,time.Duration(10)*time.Second)
	defer cancel()

	return u.repo.ListOrders(ctx, order)
}

func (u OrderServiceClient) GetOrder(ctx context.Context, orderID *genproto.OrderGetReq) (*genproto.OrderGetRes, error) {
	ctx, cancel := context.WithTimeout(ctx,time.Duration(10)*time.Second)
	defer cancel()

	return u.repo.GetOrder(ctx, orderID)
}
