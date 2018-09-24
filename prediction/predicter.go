package prediction

func (p *Predict) Predict() (float64, float64) {

	norVal := (p.Km - p.Minval) / (p.Maxval - p.Minval)

	predict := p.Teta0 + p.Teta1*norVal

	return predict, p.Dvi
}
