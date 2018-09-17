package prediction

import (
	"strconv"
)

type Norm interface {
	Normalize(km string) *Predict
}

func (p *Predict) Normalize(km string) *Predict {
	var err error

	p.Km, err = strconv.ParseFloat(km, 64)
	ErrorHandler(err)

	return p
}
