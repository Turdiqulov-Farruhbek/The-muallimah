package repository

import (
	"context"

	"gitlab.com/acumen5524834/shop-service/internal/entity"
)

type Books interface {
	CreateBook(ctx context.Context, book *entity.BookCreate) error
	UpdateBook(ctx context.Context, book *entity.BookUpdate) error
	DeleteBook(ctx context.Context, bookID string) error
	ListBooks(ctx context.Context, book *entity.BookFilter) (*entity.BookList, error)
	GetBook(ctx context.Context, bookID string) (*entity.BookGet, error)
	AddPicture(ctx context.Context, picture *entity.BookPicture) error
	DeletePicture(ctx context.Context, picture *entity.BookPicture) error
}
type Posts interface {
	CreatePost(ctx context.Context,post *entity.PostCreate) error
	GetPost(ctx context.Context,postID string) (*entity.PostGet,error)
	UpdatePost(ctx context.Context,post *entity.PostUpdate) error
	DeletePost(ctx context.Context,postID string) error
	GetPosts(ctx context.Context,post *entity.Pagination) (*entity.PostList,error)
	AddPostPicture(ctx context.Context,post *entity.PostPicture) error
	DeletePostPicture(ctx context.Context,post *entity.PostPicture) error
}

type Products interface {
	CreateProduct(ctx context.Context,product *entity.ProductCreate) error
	UpdateProduct(ctx context.Context,product *entity.ProductUpdate) error
	DeleteProduct(ctx context.Context,productID string) error
	ListProducts(ctx context.Context,product *entity.ProductFilter) (*entity.ProductList,error)
	GetProduct(ctx context.Context,productID string) (*entity.ProductGet,error)
	AddPicture(ctx context.Context,product *entity.ProductPicture) error
	DeletePicture(ctx context.Context,product *entity.ProductPicture) error
}
