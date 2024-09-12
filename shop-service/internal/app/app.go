package app

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gitlab.com/acumen5524834/shop-service/internal/delivery/grpc/server"
	"gitlab.com/acumen5524834/shop-service/internal/delivery/grpc/services"
	mdb "gitlab.com/acumen5524834/shop-service/internal/infrastructure/repository/mongo"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/config"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/genproto"
	mongosh "gitlab.com/acumen5524834/shop-service/internal/pkg/mongo"
	usecase "gitlab.com/acumen5524834/shop-service/internal/usecase/service"
	"go.mongodb.org/mongo-driver/mongo"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	Config     *config.Config
	Logger     *zap.Logger
	DB         *mongosh.MongoDB
	Postg      *sql.DB
	GrpcServer *grpc.Server
}

func NewApp(cfg *config.Config) (*App, error) {
	// Initialize MongoDB connection
	mongoDB, err := mongosh.NewMongoDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Initialize PostgreSQL connection
	pgm, err := mongosh.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	// Initialize gRPC server with middleware
	grpcServer := grpc.NewServer()

	return &App{
		Config:     cfg,
		Logger:     logger,
		DB:         mongoDB,
		Postg:      pgm.DB,
		GrpcServer: grpcServer,
	}, nil
}

func (a *App) Run() error {
	if a.Config == nil {
		return fmt.Errorf("config is nil")
	}

	if a.DB == nil {
		return fmt.Errorf("mongo database connection is nil")
	}

	if a.Postg == nil {
		return fmt.Errorf("postgres database connection is nil")
	}

	// Context timeout initialization
	contextTimeout, err := time.ParseDuration(a.Config.Context.Timeout)
	if err != nil {
		return fmt.Errorf("error during parse duration for context timeout: %w", err)
	}

	// Repository initialization
	bookRepo := mdb.NewBookManager(a.DB.Database)
	postRepo := mdb.NewPostManager(a.DB.Database)
	productRepo := mdb.NewProductManager(a.DB.Database)

	// UseCase initialization
	bookUseCase := usecase.NewBookService(contextTimeout, bookRepo)
	postUseCase := usecase.NewPostService(contextTimeout, postRepo)
	productUseCase := usecase.NewProductService(contextTimeout, productRepo)

	// Register gRPC services
	genproto.RegisterBookServiceServer(a.GrpcServer, services.NewRPC(a.Logger, bookUseCase))
	genproto.RegisterPostServiceServer(a.GrpcServer, services.NewPostRPC(a.Logger, postUseCase))
	genproto.RegisterProductServiceServer(a.GrpcServer, services.NewProductRPC(a.Logger, productUseCase))
	genproto.RegisterOrderServiceServer(a.GrpcServer, services.NewOrderServiceClient(mdb.NewOrderRepository(a.Postg, &mongo.Client{})))

	// Start gRPC server
	a.Logger.Info("gRPC Server Listening", zap.String("url", a.Config.RPCPort))
	if err := server.Run(a.Config, a.GrpcServer); err != nil {
		return fmt.Errorf("gRPC failed to serve over %s: %w", a.Config.RPCPort, err)
	}

	return nil
}

func (a *App) Stop() {
	// Stop gRPC server
	a.GrpcServer.Stop()

	// Close MongoDB connection
	a.DB.Close()

	// Sync logger
	if err := a.Logger.Sync(); err != nil {
		log.Printf("failed to sync logger: %v", err)
	}

	// Close PostgreSQL connection
	if err := a.Postg.Close(); err != nil {
		log.Printf("failed to close PostgreSQL connection: %v", err)
	}
}
