package model

import (
	"strconv"
	"fmt"
)

type Norm interface {
	Normalize(data *Data, km string)
}

func (p *Predict) Normalize(data *Data, km string) error {
	var err error

	p.Km, err = strconv.ParseFloat(km, 64)
	if err != nil {
		fmt.Println(err)
	}

	p.Teth0 = data.Teth0
	p.Teth1 = data.Teth1
	p.Dvi = data.Dvi
	p.Max = data.Max
	p.Min = data.Min

	return nil
}
