package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 564
	var aa float64 = math.Pow(a, 2)
	var b float64 = 684
	var bb float64 = math.Pow(b, 2)
	var c float64 = 63
	var cc float64 = math.Pow(c, 2)
	var summ float64 = (aa + bb + cc)
	var sub float64 = (aa - bb - cc)
	var mult float64 = (aa * bb * cc)
	var div float64 = (aa / bb / cc)
	fmt.Println(summ, sub, mult, div)

}
