package payment

import (
	"OmiseChallenges/app/domain/model"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type ChargeAPI struct {
	client *omise.Client
}

const (
	// Read these from environment variables or configuration files!
	OmisePublicKey = "pkey_test_521w1g1t7w4x4rd22z0"
	OmiseSecretKey = "skey_test_521w1g1t6yh7sx4pu8n"
)

func NewChargeAPI() *ChargeAPI {

	client, err := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if err != nil {
		return nil
	}

	return &ChargeAPI{
		client: client,
	}
}

func (r *ChargeAPI) CreateToken(tumboon model.Tumboon) (*omise.Token, *operations.CreateToken, error) {
	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            tumboon.Name,
		Number:          tumboon.CCNumber,
		ExpirationMonth: tumboon.ExpMonth,
		ExpirationYear:  tumboon.ExpYear,
	}
	if err := r.client.Do(token, createToken); err != nil {
		return nil, nil, err
	}
	return token, createToken, nil
}

func (r *ChargeAPI) DoCharge(token *omise.Token, amount int64) (*omise.Charge, *operations.CreateCharge, error) {
	// Creates a charge from the token
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   amount, // ฿ 1,000.00
		Currency: "thb",
		Card:     token.ID,
	}
	if err := r.client.Do(charge, createCharge); err != nil {
		return nil, nil, err
	}
	return charge, createCharge, nil
}