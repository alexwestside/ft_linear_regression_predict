package model

import (
	"fmt"
	"math"
)

const dvi = 142.19110308065717

type Printer interface {
	Print(res float64)
}


func (p *Predict) Print(res float64) {

	fmt.Println(fmt.Sprintf("Price is %.2f +/- %.2f", math.Ceil(res*100)/100, math.Ceil(dvi*100)/100))

}
