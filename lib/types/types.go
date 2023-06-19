package types

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type HostUpTime struct {
	Days    uint64
	Hours   uint64
	Minutes uint64
}

type AverageLoad struct {
	One     float64
	Five    float64
	Fifteen float64
}

type PercentMemoryUsed struct {
	PMU int
}

type Metrics struct {
	HUT HostUpTime
	AL  AverageLoad
	PMU PercentMemoryUsed
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	Number   int    `json:"number"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	Number int    `json:"number"`
}
type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Number            int       `json:"number"`
	Balance           int64     `json:"balance"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"createdAt"`
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type APIError struct {
	Error string `json:"error"`
}

func NewAccount(firstName, lastName, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encpw),
		Number:            (rand.Intn(1000000)),
		CreatedAt:         time.Now().UTC(),
	}, nil
}

func (a *Account) ValidatePassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}
