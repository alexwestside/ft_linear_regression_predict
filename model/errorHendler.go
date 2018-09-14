package model

import (
	"fmt"
	"os"
)

type ErrorsHandler interface {
	ErrorHandler(err error)
}

func (l *Predict) ErrorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}