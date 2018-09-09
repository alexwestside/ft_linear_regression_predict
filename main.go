package main

import (
	"github.com/ft_linear_regression_predict/predict"
	"fmt"
)

const (
	path = "data.yaml"
	km   = "176000"
)

func main() {

	prd := predict.New()

	data, err := prd.Reader(path)
	if err != nil {
		fmt.Println(err)
	}

	err = prd.Normalizer(data, km)
	if err != nil {
		fmt.Println(err)
	}

	res := prd.TryPredict()

	prd.Printer(res)
}
