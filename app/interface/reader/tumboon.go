package reader

import (
	"OmiseChallenges/app/domain/model"
	"OmiseChallenges/pkg/cipher"
	"fmt"
	"github.com/jszwec/csvutil"
	"os"
)

type tumboonRepository struct {
	reader *cipher.Rot128Reader
}

func NewTumboonRepository(file *os.File) *tumboonRepository {

	reader, err := cipher.NewRot128Reader(file)

	if err != nil {
		return nil
	}

	return &tumboonRepository{
		reader: reader,
	}
}

func (r *tumboonRepository) readBytes() []byte {
	return r.reader.DecodeToBytes()
}

func (r *tumboonRepository) GetTumboonList() (*[]model.Tumboon, error) {

	var tumboons []model.Tumboon
	bytes := r.readBytes()

	//err := csvmapper.Map(bytes, tumboons)
	if err := csvutil.Unmarshal(bytes, &tumboons); err != nil {
		fmt.Println("error:", err)
	}

	//if err!= nil {
	//	return nil, err
	//}

	return &tumboons, nil
}