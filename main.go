package main

import (
	"github.com/ft_linear_regression_predict/model"
	"os"
)

func main() {
	prd := model.New()

	err := prd.NewCommand().Execute()
	prd.ErrorHandler(err)

	os.Exit(0)
}
