package repository

import "OmiseChallenges/app/domain/model"

type TumboonReaderRepository interface {
	GetTumboonList() (*[]model.Tumboon, error)
}
