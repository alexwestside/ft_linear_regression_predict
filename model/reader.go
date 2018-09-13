package model

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
	Read(path string) (*Data, error)
}

func (p *Predict) Read(path string) (*Data, error) {
	blob, errReadFile := ioutil.ReadFile(path)
	if errReadFile != nil {
		return nil, errReadFile
	}

	data := Data{}
	errYaml := yaml.Unmarshal(blob, &data)
	if errYaml != nil {
		return nil, errReadFile
	}
	return &data, nil
}
