package main

import (
	"auth-service/internal/http/handlers"
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/postgres"
	"auth-service/internal/usecases/service"
	"fmt"
)

func main() {
	cf := config.Load()
	em := config.NewErrorManager()

	db, err := postgres.ConnectDB(&cf)
	em.CheckErr(err)

	pgsql, err := managers.New(db)
	em.CheckErr(err)
	defer pgsql.Close()

	go delivery.Server(&cf, pgsql)
	em.CheckErr(err)

	us := service.NewUserService(pgsql)

	handler := handlers.NewHandler(us)
	roter := api.NewRouter(handler)

	fmt.Println("Server is running on port ", cf.AUTH_HTTP_PORT)
	if err := roter.Run(cf.AUTH_HTTP_PORT); err != nil {
		panic(err)
	}

}
