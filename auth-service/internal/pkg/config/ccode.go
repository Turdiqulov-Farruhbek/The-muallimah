package config

import (
	rdb "auth-service/internal/pkg/redis"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	l "github.com/azizbek-qodirov/logger"
	"gopkg.in/gomail.v2"
)

func SendConfirmationCode(email string, logger *l.Logger) error {
	rdb, err := rdb.NewRedisClient(context.Background())
	if err != nil {
		return err
	}
	defer rdb.Close()

	code, err := generateConfirmationCode()
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", Load().SENDER_EMAIL)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Recovery Code")
	m.SetBody("text/plain", fmt.Sprintf("Your password recovery code is: %d", code))

	d := gomail.NewDialer("smtp.gmail.com", 587, Load().SENDER_EMAIL, Load().APP_PASSWORD)

	if err := d.DialAndSend(m); err != nil {
		logger.ERROR.Println("Error sending confirmation code: ", err.Error())
		return err
	}

	err = rdb.Set(context.Background(), email, code, 3*time.Minute).Err()
	if err != nil {
		logger.ERROR.Println("Error storing confirmation code in Redis: ", err.Error())
		return fmt.Errorf("server error storing confirmation code in Redis")
	}
	return nil
}

func generateConfirmationCode() (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}
