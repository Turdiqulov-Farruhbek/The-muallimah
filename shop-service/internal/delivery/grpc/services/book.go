package services

import (
	"context"
	"fmt"
	"log/slog"

	"gitlab.com/acumen5524834/shop-service/internal/entity"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/genproto"
	"gitlab.com/acumen5524834/shop-service/internal/usecase/service"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookRPC struct {
	genproto.UnimplementedBookServiceServer
	logger      *zap.Logger
	bookUseCase usecase.Books
}

func NewRPC(logger *zap.Logger, bookUseCase usecase.Books) genproto.BookServiceServer {
	return &BookRPC{
		logger:      logger,
		bookUseCase: bookUseCase,
	}
}

func (s BookRPC) CreateBook(ctx context.Context, in *genproto.BookCreate) (*genproto.Void, error) {
	book := s.bookUseCase.CreateBook(ctx, &entity.BookCreate{
		Title:       in.Title,
		Description: in.Description,
		Price:       in.Price,
		PictureUrls: in.PictureUrls,
	})

	slog.Info("Book created", book)

	return nil, ctx.Err()
}

func (s BookRPC) UpdateBook(ctx context.Context, in *genproto.BookUpdate) (*genproto.Void, error) {
	book := s.bookUseCase.UpdateBook(ctx, &entity.BookUpdate{
		Id: in.Id,
		Body: entity.BookUpt{Title: in.Body.Title,
			Description: in.Body.Description,
			Price:       in.Body.Price},
	})
	slog.Info("Book updated", book)
	return nil, ctx.Err()
}

func (s BookRPC) DeleteBook(ctx context.Context, in *genproto.ById) (*genproto.Void, error) {
	if err := s.bookUseCase.DeleteBook(ctx, in.Id); err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	slog.Info("Book deleted")
	return nil, ctx.Err()
}
func (s BookRPC) DeletePicture(ctx context.Context, in *genproto.BookPicture) (*genproto.Void, error) {

	err := s.bookUseCase.DeletePicture(ctx, &entity.BookPicture{
		BookId:     in.BookId,
		PictureUrl: in.PictureUrl,
	})

	if err != nil {
		if err.Error() == "invalid ObjectID format" {
			s.logger.Error("Invalid ObjectID format", zap.String("bookId", in.BookId))
			return nil, status.Errorf(codes.InvalidArgument, "invalid ObjectID format for BookId: %s", in.BookId)
		}
		return nil, status.Errorf(codes.Unknown, "failed to delete book picture: %v", err)
	}

	return &genproto.Void{}, nil
}


func (s BookRPC) GetBook(ctx context.Context, in *genproto.ById) (*genproto.BookGet, error) {
	book, err := s.bookUseCase.GetBook(ctx, in.Id)

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &genproto.BookGet{
		Id:          book.Id,
		Title:       book.Title,
		Description: book.Description,
		Price:       book.Price,
		PictureUrls: book.PictureUrls,
		CreatedAt:   book.CreatedAt,
	}, nil
}

func (s BookRPC) AddPicture(ctx context.Context, in *genproto.BookPicture) (*genproto.Void, error) {
	book := s.bookUseCase.AddPicture(ctx, &entity.BookPicture{
		BookId:     in.BookId,
		PictureUrl: in.PictureUrl,
	})

	slog.Info("Book picture created", book)

	return nil, ctx.Err()
}
func (s BookRPC) ListBooks(ctx context.Context, in *genproto.BookFilter) (*genproto.BookList, error) {
	if in == nil {
		return nil, fmt.Errorf("request is nil")
	}

	books, err := s.bookUseCase.ListBooks(ctx, &entity.BookFilter{
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

	response := &genproto.BookList{
		Books: []*genproto.BookGet{},
	}

	for _, book := range books.Books {
		response.Books = append(response.Books, &genproto.BookGet{
			Id:          book.Id,
			Title:       book.Title,
			Price:       book.Price,
			PictureUrls: book.PictureUrls,
			CreatedAt:   book.CreatedAt,
		})
	}
	response.TotalCount = books.TotalCount

	return response, nil
}
