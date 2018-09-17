package prediction

type Predicter interface {
	Predict() float64
}

func (p *Predict) Predict() (float64, float64) {

	norVal := (p.Km - p.Minval) / (p.Maxval - p.Minval)

	predict := p.Teta0 + p.Teta1*norVal

	return predict, p.Dvi
}
