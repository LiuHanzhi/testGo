package main

import (
	"fmt"
	"math/big"
)

func main() {
	mul := big.NewInt(0).Mul(big.NewInt(1), big.NewInt(2))
	fmt.Println(mul)
}
