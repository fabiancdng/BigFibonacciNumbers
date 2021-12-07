/*
     Calculator for very large Fibonacci numbers

        Implemented in Go using 'math/big'

	    Copyright (c) 2021 Fabian Reinders

*/
package main

import (
	"fmt"
	"math/big"
)

func calculate(n int) *big.Int {
	numberA := big.NewInt(0)
	numberB := big.NewInt(1)

	var result *big.Int

	if n == 0 {
		return numberA
	}

	if n == 1 {
		return numberB
	}

	for i := 2; i <= n; i++ {
		result = big.NewInt(0)
		result.Add(numberA, numberB)

		numberA = numberB
		numberB = result
	}

	return result
}

func main() {
	var n int
	fmt.Print("n=")
	fmt.Scanf("%d", &n)

	result := calculate(n)

	fmt.Println("Result: ", result)
}
