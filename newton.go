package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
)

/* 
Newton's method:
    z = z - (z^2 - x)/(2z)
approaches sqrt(z)
*/

func Sqrt(x float64) float64 {
	z := 1.0
	prev := 0.0
	epsilon := 1E-7
	for math.Abs(z-prev) >= epsilon {
		prev = z
		z = z - (math.Pow(z, 2)-x)/(2*z)
		fmt.Printf("Current z: %f\n", z)
	}
	return z
}

func main() {
	var operand_p *float64 = new(float64)
	flag.Float64Var(operand_p, "operand", math.NaN(), "sqrt")
	flag.Parse()
	if math.IsNaN(*operand_p) {
		var err error
		*operand_p, err = strconv.ParseFloat(flag.Arg(0), 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Newton: %f\n", Sqrt(*operand_p))
	fmt.Printf("math.Sqrt: %f\n", math.Sqrt(*operand_p))
}
