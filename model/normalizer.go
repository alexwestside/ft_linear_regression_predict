package model

import "strconv"

type Norm interface {
	Normaliz(data *Data, km string)
}

func (p *Predict) Normaliz(data *Data, km string) error {
	var err error

	p.Teth0, err = strconv.ParseFloat(data.Teth0, 64)
	if err != nil {
		return err
	}

	p.Teth1, err = strconv.ParseFloat(data.Teth1, 64)
	if err != nil {
		return err
	}

	p.Max, err = strconv.ParseFloat(data.Max, 64)
	if err != nil {
		return err
	}

	p.Min, err = strconv.ParseFloat(data.Min, 64)
	if err != nil {
		return err
	}

	p.Km, err = strconv.ParseFloat(km, 64)
	if err != nil {
		return err
	}

	return nil
}
