package usecase

import (
	"context"
	"time"

	"gitlab.com/acumen5524834/shop-service/internal/entity"
	"gitlab.com/acumen5524834/shop-service/internal/infrastructure/repository"
)

type Posts interface {
	CreatePost(ctx context.Context, post *entity.PostCreate) error
	GetPost(ctx context.Context, postID string) (*entity.PostGet, error)
	UpdatePost(ctx context.Context, post *entity.PostUpdate) error
	DeletePost(ctx context.Context, postID string) error
	GetPosts(ctx context.Context, post *entity.Pagination) (*entity.PostList, error)
	AddPostPicture(ctx context.Context, post *entity.PostPicture) error
	DeletePostPicture(ctx context.Context, post *entity.PostPicture) error
}
type postService struct {
	repo       repository.Posts
	ctxTimeout time.Duration
}

func NewPostService(ctxTimeout time.Duration, repo repository.Posts) Posts {
	return &postService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u postService) CreatePost(ctx context.Context, post *entity.PostCreate) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.CreatePost(ctx, post)
}

func (u postService) UpdatePost(ctx context.Context, post *entity.PostUpdate) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.UpdatePost(ctx, post)
}

func (u postService) DeletePost(ctx context.Context, postID string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.DeletePost(ctx, postID)
}

func (u postService) GetPosts(ctx context.Context, post *entity.Pagination) (*entity.PostList, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.GetPosts(ctx, post)
}

func (u postService) GetPost(ctx context.Context, postID string) (*entity.PostGet, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.GetPost(ctx, postID)
}

func (u postService) AddPostPicture(ctx context.Context, picture *entity.PostPicture) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.AddPostPicture(ctx, picture)
}

func (u postService) DeletePostPicture(ctx context.Context, picture *entity.PostPicture) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.DeletePostPicture(ctx, picture)
}
