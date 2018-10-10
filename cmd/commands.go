package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ft_linear_regression_predict/prediction"
	"fmt"
	"math"
)

func Commands(rootCmd *cobra.Command) *cobra.Command {
	rootCmd.PersistentFlags().String("t", "", "read target")
	rootCmd.PersistentFlags().String("i", "", "read independent variable")
	rootCmd.AddCommand(predict())
	return rootCmd
}

func predict() (cmd *cobra.Command) {
	return &cobra.Command{
		Use:   "predict",
		Short: "predict",
		Run: func(cmd *cobra.Command, args []string) {
			var path string
			if path = cmd.Flag("t").Value.String(); path == "" {
				handler("target")
			} else if iv := cmd.Flag("i").Value.String(); iv == "" {
				handler("independent variable")
			} else {
				result, dvi := prediction.New().Read(path, iv).Predict()
				output(result, dvi)
				//fmt.Println(fmt.Sprintf("RESULT: Price is %.2f +/- %.2f", math.Ceil(result*100)/100, math.Ceil(dvi*100)/100))
			}
		},
	}
}


func output(result float64, dvi float64) {
	fmt.Println("------------------------------------------------------------------")
	fmt.Print("Liner regression -> PREDICT MODE")
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(fmt.Sprintf("RESULT: Price is %.2f +/- %.2f", math.Ceil(result*100)/100, math.Ceil(dvi*100)/100))
	fmt.Println("------------------------------------------------------------------")
}