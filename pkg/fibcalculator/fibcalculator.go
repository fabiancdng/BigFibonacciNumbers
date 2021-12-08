package fibcalculator

import "math/big"

// Calculates the nth fibonacci number.
func FibCalculate(n int) *big.Int {
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
