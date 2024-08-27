package usecase

import (
	"context"
	"time"

	"gitlab.com/acumen5524834/shop-service/internal/entity"
	"gitlab.com/acumen5524834/shop-service/internal/infrastructure/repository"
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
type bookService struct {
	BaseUseCase
	repo       repository.Books
	ctxTimeout time.Duration
}

func NewBookService(ctxTimeout time.Duration, repo repository.Books) Books {
	return &bookService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u bookService) CreateBook(ctx context.Context, book *entity.BookCreate)  error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.CreateBook(ctx, book)
}

func (u bookService) UpdateBook(ctx context.Context, book *entity.BookUpdate) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.UpdateBook(ctx, book)
}

func (u bookService) DeleteBook(ctx context.Context, bookID string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.DeleteBook(ctx, bookID)
}

func (u bookService) ListBooks(ctx context.Context, book *entity.BookFilter) (*entity.BookList,error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.ListBooks(ctx, book)
}

func (u bookService) GetBook(ctx context.Context, bookID string) (*entity.BookGet, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.GetBook(ctx, bookID)
}

func (u bookService) AddPicture(ctx context.Context, picture *entity.BookPicture) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.AddPicture(ctx, picture)
}

func (u bookService) DeletePicture(ctx context.Context, picture *entity.BookPicture) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.DeletePicture(ctx, picture)
}
