package main

import (
	"github.com/ft_linear_regression_predict/model"
	"fmt"
)

const (
	path = "../ft_linear_regression_train/data.yaml"
	km   = "176000"
)

func main() {

	prd := model.New()

	data, err := prd.Read(path)
	if err != nil {
		fmt.Println(err)
	}

	err = prd.Normalize(data, km)
	if err != nil {
		fmt.Println(err)
	}

	prd.Print(prd.Predict())
}
