package usecase

import (
	"OmiseChallenges/app/domain/repository"
	"OmiseChallenges/app/domain/service"
	"OmiseChallenges/app/interface/payment"
	"OmiseChallenges/app/interface/reader"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type tumboonUseCase struct {
	readerRepo repository.TumboonReaderRepository
	paymentRepo repository.TumboonPaymentRepository
	service *service.TumboonPaymentService
}

func BuildTumboonUseCase(file *os.File) (*tumboonUseCase, error){
	readerRepo := reader.NewTumboonRepository(file)
	if readerRepo == nil {
		return nil, errors.New("Couldn't Build Tumbon Usecase. File may be error or not exist.")
	}

	paymentRepo := payment.NewChargeAPI()
	if paymentRepo == nil {
		return nil, errors.New("Couldn't Build Tumbon Usecase. Payment Gateway Public/Secret may be error or not exist.")
	}

	service := service.NewTumboonPaymentService(paymentRepo)

	return NewTumboonUseCase(readerRepo, paymentRepo, service), nil
}

func NewTumboonUseCase(readerRepo repository.TumboonReaderRepository, paymentRepo repository.TumboonPaymentRepository, service *service.TumboonPaymentService) *tumboonUseCase {
	return &tumboonUseCase{
		readerRepo: readerRepo,
		paymentRepo: paymentRepo,
		service: service,
	}
}

func (u *tumboonUseCase) DoTumboonFromList() error {

	tumboons, err := u.readerRepo.GetTumboonList()

	if err != nil {
		return err
	}

	fmt.Println(tumboons)

	return nil
}