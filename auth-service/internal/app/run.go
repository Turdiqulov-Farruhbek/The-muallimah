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

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(cf config.Config) {
	em := config.NewErrorManager()

	db, err := postgres.ConnectDB(&cf)
	em.CheckErr(err)

	pgsql, err := repo.New(db)
	em.CheckErr(err)
	defer pgsql.Close()

	go service.Server(&cf, pgsql)
	em.CheckErr(err)

	AuthConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", cf.AUTH_GRPC_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	em.CheckErr(err)
	defer AuthConn.Close()

	handler := handlers.NewHandler(AuthConn)
	roter := api.NewRouter(handler)

	fmt.Println("Server is running on port:", cf.AUTH_HTTP_PORT)
	if err := roter.Run(cf.AUTH_HTTP_PORT); err != nil {
		panic(err)
	}
}
