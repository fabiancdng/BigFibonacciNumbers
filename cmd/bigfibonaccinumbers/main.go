/*
     Calculator for very large Fibonacci numbers

        Implemented in Go using 'math/big'

	    Copyright (c) 2021 Fabian Reinders

*/
package main

import (
	"fmt"

	"github.com/fabiancdng/BigFibonacciNumbers/pkg/fibcalculator"
)

func main() {
	var n int
	fmt.Print("n=")
	fmt.Scanf("%d", &n)

	result := fibcalculator.FibCalculate(n)

	fmt.Println("Result: ", result)
}
