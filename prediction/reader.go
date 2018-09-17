package prediction

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Data struct {
	Teth0 float64
	Teth1 float64
	Min   float64
	Max   float64
	Dvi   float64
}

type Reader interface {
	Read(path string) *Data
}

func (p *Predict) Read(path string) *Predict{
	blob, errReadFile := ioutil.ReadFile(path)
	ErrorHandler(errReadFile)

	errYaml := yaml.Unmarshal(blob, &p)
	ErrorHandler(errYaml)

	return p
}
