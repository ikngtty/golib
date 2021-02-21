package simple

import "fmt"

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Min returns the minimum value of the specified values.
func Min(values ...int) int {
	if len(values) == 0 {
		panic("no values")
	}

	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum value of the specified values.
func Max(values ...int) int {
	if len(values) == 0 {
		panic("no values")
	}

	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

// Pow returns base^exponent.
func Pow(base, exponent int) int {
	if exponent < 0 {
		panic(fmt.Sprintf("invalid exponent: %d", exponent))
	}

	answer := 1
	for i := 0; i < exponent; i++ {
		answer *= base
	}
	return answer
}

// Ceil returns ceil(divident/dividor).
func Ceil(divident, dividor int) int {
	if dividor == 0 {
		panic("divide by zero")
	}

	quo := divident / dividor
	rem := divident % dividor

	if rem != 0 {
		if (divident > 0 && dividor > 0) ||
			(divident < 0 && dividor < 0) {
			return quo + 1
		}
	}
	return quo
}
