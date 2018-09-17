package prediction

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strconv"
	"fmt"
)


type Reader interface {
	Read(path string) *Predict
}

func (p *Predict) Read(path string, iv string) *Predict{
	blob, errReadFile := ioutil.ReadFile(path)
	ErrorHandler(errReadFile)

	errYaml := yaml.Unmarshal(blob, &p)
	ErrorHandler(errYaml)

	var err error
	p.Km, err = strconv.ParseFloat(iv, 64)
	ErrorHandler(err)

	fmt.Println(p)

	return p
}
