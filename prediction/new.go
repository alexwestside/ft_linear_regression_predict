package prediction

type Predict struct {
	Teta0 float64
	Teta1 float64
	Minval   float64
	Maxval   float64
	Km    float64
	Dvi   float64
}

func New() *Predict {
	return &Predict{}
}

