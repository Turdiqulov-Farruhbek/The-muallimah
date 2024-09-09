package mongo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gitlab.com/acumen5524834/shop-service/internal/pkg/config"
)

type Postgres struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*Postgres, error) {
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Database)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}
func (db *Postgres) Close() error {
	return db.DB.Close()
}
