package services

import (
	"context"
	"log/slog"

	"gitlab.com/acumen5524834/shop-service/internal/entity"
	"gitlab.com/acumen5524834/shop-service/internal/infrastructure/grpc_service_clients"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/genproto"
	"gitlab.com/acumen5524834/shop-service/internal/usecase"
	"go.uber.org/zap"
)

type postRPC struct {
	genproto.UnimplementedPostServiceServer
	logger      *zap.Logger
	postUseCase usecase.Posts
	client      grpc_service_clients.ServiceClients
}

func NewPostRPC(logger *zap.Logger, postUseCase usecase.Posts, client *grpc_service_clients.ServiceClients) genproto.PostServiceServer {
	return &postRPC{
		logger:      logger,
		postUseCase: postUseCase,
		client:      *client,
	}
}

func (s postRPC) CreatePost(ctx context.Context, in *genproto.PostCreate) (*genproto.Void, error) {
	post := s.postUseCase.CreatePost(ctx, &entity.PostCreate{
		Title:       in.Title,
		Content:     in.Content,
		PictureUrls: in.PictureUrls,
	})

	slog.Info("post created", post)

	return nil, ctx.Err()
}

func (s postRPC) UpdatePost(ctx context.Context, in *genproto.PostUpdate) (*genproto.Void, error) {
	post := s.postUseCase.UpdatePost(ctx, &entity.PostUpdate{
		ID: in.Id,
		Body: entity.PostUpt{
			Title:   in.Body.Title,
			Content: in.Body.Content},
	})
	slog.Info("post updated", post)
	return nil, ctx.Err()
}

func (s postRPC) DeletePost(ctx context.Context, in *genproto.ById) (*genproto.Void, error) {
	if err := s.postUseCase.DeletePost(ctx, in.Id); err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	slog.Info("post deleted")
	return nil, ctx.Err()
}

func (s postRPC) DeletePostPicture(ctx context.Context, in *genproto.PostPicture) (*genproto.Void, error) {
	s.logger.Info("DeletePostPicture called", zap.String("PostId", in.PostId), zap.String("PictureUrl", in.PictureUrl))

	err := s.postUseCase.DeletePostPicture(ctx, &entity.PostPicture{
		PostID:     in.PostId,
		PictureUrl: in.PictureUrl,
	})

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	s.logger.Info("Post picture deleted successfully", zap.String("postId", in.PostId), zap.String("pictureUrl", in.PictureUrl))
	return &genproto.Void{}, nil
}

func (s postRPC) GetPost(ctx context.Context, in *genproto.ById) (*genproto.PostGet, error) {
	post, err := s.postUseCase.GetPost(ctx, in.Id)

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &genproto.PostGet{
		Id:          post.ID,
		Title:       post.Title,
		Content:     post.Content,
		PictureUrls: post.PictureUrls,
		CreatedAt:   post.CreatedAt,
	}, nil
}

func (s postRPC) AddPostPicture(ctx context.Context, in *genproto.PostPicture) (*genproto.Void, error) {
	post := s.postUseCase.AddPostPicture(ctx, &entity.PostPicture{
		PostID:     in.PostId,
		PictureUrl: in.PictureUrl,
	})

	slog.Info("post picture created", post)

	return nil, ctx.Err()
}
func (s postRPC) GetPosts(ctx context.Context,in *genproto.Pagination) (*genproto.PostList,error) {

	posts, err := s.postUseCase.GetPosts(context.Background(), &entity.Pagination{
		Limit:  int32(in.Limit),
		Offset: int32(in.Offset),
	})

	if err != nil {
		s.logger.Error("Error getting posts", zap.Error(err))
		return nil,err
	}

	postList := &genproto.PostList{
		Posts: []*genproto.PostGet{},
	}

	for _, post := range posts.Posts {
		postList.Posts = append(postList.Posts, &genproto.PostGet{
			Id:          post.ID,
			Title:       post.Title,
			Content:     post.Content,
			PictureUrls: post.PictureUrls,
			CreatedAt:   post.CreatedAt,
		})
	}
	postList.TotalCount = posts.TotalCount

	return nil,err
}
