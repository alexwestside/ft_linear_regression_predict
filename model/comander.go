package model

import (
	"github.com/spf13/cobra"
	"fmt"
)

const (
	targetFlag      = "target"
	targetFlagShort = "t"
	dataFlag        = "data"
	dataFlagShort   = "d"
)

type Commander interface {
	NewCommand() error
}

func (p *Predict) NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "PROJECT: ft_linear_regression (predict model)",
		Short: "The Linear Regression predict service",
		Long:  "Predict service",
		Run:   p.serve,
	}

	flags := cmd.Flags()
	flags.StringP(targetFlag, targetFlagShort, "", "File target with data-set")
	flags.StringP(dataFlag, dataFlagShort, "", "Target folder for write result data")

	return
}

func (p *Predict) serve(cmd *cobra.Command, args []string) {
	fmt.Println(cmd.Use)
	fmt.Println(cmd.Short, "\n")

	var df *Data
	var err error

	if path := cmd.Flag(targetFlag).Value.String(); path != "" {
		df, err = p.Read(path)
		p.ErrorHandler(err)
		fmt.Println(" -> Read data from:", path)

		if km := cmd.Flag(dataFlag).Value.String(); km != "" {
			err = p.Normalize(df, km)
			p.ErrorHandler(err)

			p.Print(p.Predict())

		} else {
			fmt.Println("Data target not defined")
		}

	} else {
		fmt.Println("Read target not defined")
	}
}
