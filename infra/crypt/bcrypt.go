package crypt

import (
	"fmt"

	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
)

var Module = fx.Module("crypt",
	fx.Provide(
		func() *Crypt {
			return NewCrypt("LOCAL") // ou "PROD-PROD"
		},
	),
)

type BCryptInterface interface {
	HashPassword(string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}

type Crypt struct {
	bcryptCost int
}

func NewCrypt(environment string) *Crypt {
	bcryptCost := bcrypt.DefaultCost
	if environment != "PROD-PROD" {
		bcryptCost = bcrypt.MinCost
	}

	return &Crypt{bcryptCost}
}

func (crypt *Crypt) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), crypt.bcryptCost)
	if err != nil {
		err = fmt.Errorf("error on generate hash password %d", err)
		return "", err
	}

	return string(bytes), nil
}

func (crypt *Crypt) CompareHashAndPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("senha incorreta")
	}
	return nil
}
