package model

import (
	"fmt"
	"math"
)

type Printer interface {
	Print(res float64)
}


func (p *Predict) Print(res float64) {

	fmt.Println("RESULT:")
	fmt.Println(fmt.Sprintf("Price is %.2f +/- %.2f", math.Ceil(res*100)/100, math.Ceil(p.Dvi*100)/100))

}
