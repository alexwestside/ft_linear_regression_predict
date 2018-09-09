package predict

import (
	"fmt"
	"math"
)

const dvi = 142.19110308065717

func (p *Predict) Printer(res float64) {

	fmt.Println(fmt.Sprintf("Price is %.2f +/- %.2f", math.Ceil(res*100)/100, math.Ceil(dvi*100)/100))

}
