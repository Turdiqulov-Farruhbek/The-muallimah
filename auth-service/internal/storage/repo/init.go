package repo

import (
	"auth-service/internal/storage"
	"database/sql"
)

type Storage struct {
	db    *sql.DB
	UserS storage.UserI
}

func New(db *sql.DB) (*Storage, error) {
	userS := NewUserManager(db)

	return &Storage{
		UserS: userS,
		db:    db,
	}, nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) User() storage.UserI {
	return s.UserS
}
