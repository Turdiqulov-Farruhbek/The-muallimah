package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gitlab.com/muallimah/notification_service/internal/pkg/config"
	"gitlab.com/muallimah/notification_service/internal/storage"
)

type Storage struct {
	db            *sql.DB
	NotificationS storage.NotificationI
}

func NewStorage(db *sql.DB,cfg *config.Config) (*Storage, error) {

	notifS := NewNotificationRepo(db, cfg)
	return &Storage{
		db:            db,
		NotificationS: notifS,
	}, nil
}
func (s *Storage) Notification() storage.NotificationI {
	return s.NotificationS
}
