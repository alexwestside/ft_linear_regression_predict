package main

import (
	"github.com/ft_linear_regression_predict/model"
	"os"
)

const (
	path = "../ft_linear_regression_train/data.yaml"
	km   = "176000"
)

func main() {
	prd := model.New()

	err := prd.NewCommand().Execute()
	prd.ErrorHandler(err)

	os.Exit(0)
}
