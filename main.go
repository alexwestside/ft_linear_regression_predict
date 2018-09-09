package main

import (
	"github.com/ft_linear_regression_predict/model"
	"fmt"
)

const (
	path = "data.yaml"
	km   = "176000"
)

func main() {

	prd := model.New()

	data, err := prd.Read(path)
	if err != nil {
		fmt.Println(err)
	}

	err = prd.Normaliz(data, km)
	if err != nil {
		fmt.Println(err)
	}

	res := prd.Predict()

	prd.Print(res)
}
