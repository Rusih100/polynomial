package main

import (
	"fmt"
	"math/big"
	"polynomial/polynomial"
)

func main() {

	// Create poly

	// X^3 + 5x - 4
	polyCoefficient := []*big.Int{
		big.NewInt(-4),
		big.NewInt(5),
		big.NewInt(0),
		big.NewInt(1),
	}

	poly := polynomial.NewPolynomial(polyCoefficient)

	fmt.Println(poly)

}
