package predict

type Model interface {
	Reader(path string) (*Data, error)
	TryPredict() float64
	Normalizer(data *Data, km string) error
	Printer(res float64)
}

type Predict struct {
	Teth0 float64
	Teth1 float64
	Min   float64
	Max   float64
	Km    float64
}

func New() *Predict {
	return &Predict{}
}
