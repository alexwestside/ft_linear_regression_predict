package model

type Predict struct {
	Teth0 float64
	Teth1 float64
	Min   float64
	Max   float64
	Km    float64
	Dvi   float64
	Reader
	Norm
	Printer
	Predicter
	Commander
}

func New() *Predict {
	return &Predict{}
}
