package postgres

import (
	"auth-service/internal/pkg/config"
	"database/sql"
	"fmt"

	l "github.com/azizbek-qodirov/logger"
	_ "github.com/lib/pq"
)

func ConnectDB(cf *config.Config, logger *l.Logger) (*sql.DB, error) {
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cf.DB_USER, cf.DB_PASSWORD, cf.DB_HOST, cf.DB_PORT, cf.DB_NAME)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.ERROR.Panicln("Postgres not connected due to error: " + err.Error())
	}
	logger.INFO.Println("Postgres connected")
	
	return db, nil
}
