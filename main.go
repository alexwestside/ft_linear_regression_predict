package main

import (
	"github.com/ft_linear_regression/ft_linear_regression_predict/cmd"
	"github.com/spf13/cobra"
)

func main() {

	cmd.Commands(&cobra.Command{
		Use:   "app",
		Short: "The Linear Regression prediction service",
		Long:  "PROJECT: ft_linear_regression (prediction)",
	}).Execute()
}
