package services

import (
	"context"
	"fmt"
	"log/slog"

	"gitlab.com/acumen5524834/shop-service/internal/entity"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/genproto"
	"gitlab.com/acumen5524834/shop-service/internal/usecase/service"
	"go.uber.org/zap"
)

type ProductRPC struct {
	genproto.UnimplementedProductServiceServer
	logger         *zap.Logger
	productUseCase usecase.Products
}

func NewProductRPC(logger *zap.Logger, productUseCase usecase.Products) genproto.ProductServiceServer {
	return &ProductRPC{
		logger:         logger,
		productUseCase: productUseCase,
	}
}

func (s ProductRPC) CreateProduct(ctx context.Context, in *genproto.ProductCreate) (*genproto.Void, error) {
	product := s.productUseCase.CreateProduct(ctx, &entity.ProductCreate{
		Title:       in.Title,
		Description: in.Description,
		Price:       in.Price,
		PictureUrls: in.PictureUrls,
	})

	slog.Info("product created", product)

	return nil, ctx.Err()
}

func (s ProductRPC) UpdateProduct(ctx context.Context, in *genproto.ProductUpdate) (*genproto.Void, error) {
	product := s.productUseCase.UpdateProduct(ctx, &entity.ProductUpdate{
		Id: in.Id,
		Body: entity.ProductUpt{
			Title:       in.Body.Title,
			Description: in.Body.Description,
			Price:       in.Body.Price},
	})
	slog.Info("product updated", product)
	return nil, ctx.Err()
}

func (s ProductRPC) DeleteProduct(ctx context.Context, in *genproto.ById) (*genproto.Void, error) {
	if err := s.productUseCase.DeleteProduct(ctx, in.Id); err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	slog.Info("product deleted")
	return nil, ctx.Err()
}

func (s ProductRPC) DeletePicture(ctx context.Context, in *genproto.ProductPicture) (*genproto.Void, error) {
	if err := s.productUseCase.DeletePicture(ctx, &entity.ProductPicture{ProductId: in.ProductId, PictureUrl: in.PictureUrl}); err != nil {
		s.logger.Error("Failed to delete product picture", zap.String("productId", in.ProductId), zap.String("pictureUrl", in.PictureUrl), zap.Error(err))
		return nil, err
	}

	slog.Info("product picture deleted")
	return nil, ctx.Err()
}

func (s ProductRPC) GetProduct(ctx context.Context, in *genproto.ById) (*genproto.ProductGet, error) {
	product, err := s.productUseCase.GetProduct(ctx, in.Id)

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &genproto.ProductGet{
		Id:          product.Id,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		PictureUrls: product.PictureUrls,
		CreatedAt:   product.CreatedAt,
	}, nil
}

func (s ProductRPC) AddPicture(ctx context.Context, in *genproto.ProductPicture) (*genproto.Void, error) {
	product := s.productUseCase.AddPicture(ctx, &entity.ProductPicture{
		ProductId:  in.ProductId,
		PictureUrl: in.PictureUrl,
	})

	slog.Info("product picture created", product)

	return nil, ctx.Err()
}
func (s ProductRPC) ListProducts(ctx context.Context, in *genproto.ProductFilter) (*genproto.ProductList, error) {
	if in == nil {
		return nil, fmt.Errorf("request is nil")
	}

	products, err := s.productUseCase.ListProducts(ctx, &entity.ProductFilter{
		Title:     in.Title,
		PriceFrom: float32(in.PriceFrom),
		PriceTo:   float32(in.PriceTo),
		Pagination: entity.Pagination{
			Limit:  int32(in.Pagination.Limit),
			Offset: int32(in.Pagination.Offset),
		},
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	response := &genproto.ProductList{
		Products: []*genproto.ProductGet{},
	}

	for _, product := range products.Products {
		response.Products = append(response.Products, &genproto.ProductGet{
			Id:          product.Id,
			Title:       product.Title,
			Price:       product.Price,
			PictureUrls: product.PictureUrls,
			CreatedAt:   product.CreatedAt,
		})
	}
	response.TotalCount = products.TotalCount

	return response, nil
}
