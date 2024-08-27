package app

import (
	api "auth-service/internal/http"
	"auth-service/internal/http/handlers"
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/postgres"
	"auth-service/internal/storage/repo"
	"auth-service/internal/usecases/service"
	"fmt"

	_ "auth-service/internal/http/docs"

	l "github.com/azizbek-qodirov/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(cf config.Config) {
	em := config.NewErrorManager()
	logger, err := l.NewLogger(&l.LogFileConfigs{
		Directory: "logs",
		Filename:  "app.log",
		Stdout:    false,
		Include:   l.DateTime | l.Loglevel | l.ShortFileName,
	})
	em.CheckErr(err)

	db, err := postgres.ConnectDB(&cf, logger)
	em.CheckErr(err)

	pgsql, err := repo.New(db)
	em.CheckErr(err)
	defer pgsql.Close()

	go service.Server(&cf, *logger, pgsql)
	em.CheckErr(err)

	AuthConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", cf.AUTH_GRPC_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	em.CheckErr(err)
	defer AuthConn.Close()

	handler := handlers.NewHandler(AuthConn, logger)
	roter := api.NewRouter(handler, logger)

	fmt.Println("Server is running on port:", cf.AUTH_HTTP_PORT)
	if err := roter.Run(cf.AUTH_HTTP_PORT); err != nil {
		panic(err)
	}
}
