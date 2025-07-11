package main

import (
	"fmt"
	"math/big"
)

func L122() {
	a := big.NewInt(1234567890123456789)
	b := big.NewInt(9876543210)
	resultMul := new(big.Int).Mul(a, b)
	resultAdd := new(big.Int).Add(a, b)
	resultSub := new(big.Int).Sub(a, b)
	resultDiv := new(big.Int).Div(a, b)
	fmt.Print(resultMul, resultAdd, resultDiv, resultSub)

}
