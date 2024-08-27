package app

import (
	"fmt"
	"time"

	"gitlab.com/acumen5524834/shop-service/internal/delivery/grpc/server"
	"gitlab.com/acumen5524834/shop-service/internal/delivery/grpc/services"
	"gitlab.com/acumen5524834/shop-service/internal/infrastructure/grpc_service_clients"
	mdb "gitlab.com/acumen5524834/shop-service/internal/infrastructure/repository/mongo"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/config"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/genproto"
	mongosh "gitlab.com/acumen5524834/shop-service/internal/pkg/mongo"
	"gitlab.com/acumen5524834/shop-service/internal/usecase"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	Config         *config.Config
	Logger         *zap.Logger
	DB             *mongosh.MongoDB
	GrpcServer     *grpc.Server
	ServiceClients grpc_service_clients.ServiceClients
}

func NewApp(cfg *config.Config) (*App, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	db, err := mongosh.NewMongoDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// grpc server init
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
			grpcctxtags.StreamServerInterceptor(),
			grpcrecovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(server.UnaryInterceptor(
			grpcmiddleware.ChainUnaryServer(
				grpcctxtags.UnaryServerInterceptor(),
				grpcrecovery.UnaryServerInterceptor(),
			),
		)),
	)

	return &App{
		Config:     cfg,
		Logger:     logger,
		DB:         db,
		GrpcServer: grpcServer,
	}, nil

}

func (a *App) Run() error {
	if a.Config == nil {
		return fmt.Errorf("config is nil")
	}

	if a.DB == nil {
		return fmt.Errorf("database connection is nil")
	}

	// context timeout initialization
	contextTimeout, err := time.ParseDuration(a.Config.Context.Timeout)
	if err != nil {
		return fmt.Errorf("error during parse duration for context timeout: %w", err)
	}

	// Initialize Service Clients
	serviceClients, err := grpc_service_clients.New(a.Config)
	if err != nil {
		return fmt.Errorf("error during initialize service clients: %w", err)
	}
	a.ServiceClients = serviceClients

	// repositories initialization
	if a.DB.Database == nil {
		return fmt.Errorf("database is nil")
	}

	bookRepo := mdb.NewBookManager(a.DB.Database)
	postRepo := mdb.NewPostManager(a.DB.Database)
	productRepo := mdb.NewProductManager(a.DB.Database)

	// useCase initialization
	bookUseCase := usecase.NewBookService(contextTimeout, bookRepo)
	postUseCase := usecase.NewPostService(contextTimeout, postRepo)
	productUseCase := usecase.NewProductService(contextTimeout, productRepo)

	genproto.RegisterBookServiceServer(a.GrpcServer, services.NewRPC(a.Logger, bookUseCase, &a.ServiceClients))
	genproto.RegisterPostServiceServer(a.GrpcServer, services.NewPostRPC(a.Logger, postUseCase, &a.ServiceClients))
	genproto.RegisterProductServiceServer(a.GrpcServer, services.NewProductRPC(a.Logger, productUseCase, &a.ServiceClients))
	a.Logger.Info("gRPC Server Listening", zap.String("url", a.Config.RPCPort))
	if err := server.Run(a.Config, a.GrpcServer); err != nil {
		return fmt.Errorf("gRPC failed to serve grpc server over %s: %w", a.Config.RPCPort, err)
	}

	return nil
}

func (a *App) Stop() {
	// closing client service connections
	a.ServiceClients.Close()
	// stop gRPC server
	a.GrpcServer.Stop()

	// database connection
	a.DB.Close()

	// zap logger sync
	err := a.Logger.Sync()
	if err != nil {
		return
	}
}
