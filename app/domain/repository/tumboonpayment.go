package repository

import (
	"OmiseChallenges/app/domain/model"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type TumboonPaymentRepository interface {
	CreateToken(tumboon model.Tumboon) (*omise.Token, *operations.CreateToken, error)
	DoCharge(token *omise.Token, amount int64) (*omise.Charge, *operations.CreateCharge, error)
}
