package csvmapper

import (
	"github.com/jszwec/csvutil"
)

func Map(data []byte, output interface{}) error {
	if err := csvutil.Unmarshal(data, &output); err != nil {
		return err
	}

	return nil
}
