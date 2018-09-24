package prediction

import (
	"strconv"
)

func (p *Predict) Normalize(km string) *Predict {
	var err error

	p.Km, err = strconv.ParseFloat(km, 64)
	ErrorHandler(err)

	return p
}
