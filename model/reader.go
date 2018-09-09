package model

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Data struct {
	Teth0 string
	Teth1 string
	Min   string
	Max   string
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
