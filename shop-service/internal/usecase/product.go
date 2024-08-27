package usecase

import (
	"context"
	"time"

	"gitlab.com/acumen5524834/shop-service/internal/entity"
	"gitlab.com/acumen5524834/shop-service/internal/infrastructure/repository"
)

type Products interface {
	CreateProduct(ctx context.Context, product *entity.ProductCreate) error
	UpdateProduct(ctx context.Context, product *entity.ProductUpdate) error
	DeleteProduct(ctx context.Context, productID string) error
	ListProducts(ctx context.Context, product *entity.ProductFilter) (*entity.ProductList, error)
	GetProduct(ctx context.Context, productID string) (*entity.ProductGet, error)
	AddPicture(ctx context.Context, product *entity.ProductPicture) error
	DeletePicture(ctx context.Context, product *entity.ProductPicture) error
}

type productService struct {
	BaseUseCase
	repo       repository.Products
	ctxTimeout time.Duration
}

func NewProductService(ctxTimeout time.Duration, repo repository.Products) Products {
	return &productService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u productService) CreateProduct(ctx context.Context, product *entity.ProductCreate) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.CreateProduct(ctx, product)
}

func (u productService) UpdateProduct(ctx context.Context, product *entity.ProductUpdate) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.UpdateProduct(ctx, product)
}

func (u productService) DeleteProduct(ctx context.Context, productID string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.DeleteProduct(ctx, productID)
}

func (u productService) ListProducts(ctx context.Context, product *entity.ProductFilter) (*entity.ProductList, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.ListProducts(ctx, product)
}

func (u productService) GetProduct(ctx context.Context, productID string) (*entity.ProductGet, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.GetProduct(ctx, productID)
}

func (u productService) AddPicture(ctx context.Context, productpic *entity.ProductPicture) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.AddPicture(ctx, productpic)
}

func (u productService) DeletePicture(ctx context.Context, productpic *entity.ProductPicture) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.DeletePicture(ctx, productpic)
}
