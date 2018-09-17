package prediction

type Predicter interface {
	Predict() float64
}

func (p *Predict) Predict() (float64, float64) {

	norVal := (p.Km - p.Min) / (p.Max - p.Min)

	predict := p.Teth0 + p.Teth1*norVal

	return predict, p.Dvi
}
