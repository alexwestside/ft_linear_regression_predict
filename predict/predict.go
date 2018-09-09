package predict


func (p *Predict) TryPredict() float64 {

	norVal := (p.Km - p.Min) / (p.Max - p.Min)

	predict := p.Teth0 + p.Teth1 * norVal

	return predict
}


