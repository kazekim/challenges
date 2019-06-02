package service

import (
	"OmiseChallenges/app/domain/repository"
)

type TumboonPaymentService struct {
	repo repository.TumboonPaymentRepository
}

func NewTumboonPaymentService(repo repository.TumboonPaymentRepository) *TumboonPaymentService {
	return &TumboonPaymentService{
		repo:repo,
	}
}
