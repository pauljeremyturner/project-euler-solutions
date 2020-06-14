package main

import "math"

func SumEvenFibonacciSimple(limit int) int {
	sum := 0
	for i, j := 1, 1; j < limit; i, j = i+j, i {
		if math.Mod(float64(i), 2.0) == 0 {
			sum += i
		}
	}
	return sum
}



func SumEvenFibonacciBitshift(limit int) int {
	sum := 0
	for i, j := 1, 1; j < limit; i, j = i+j, i {
		if even(i) {
			sum += i
		}
	}
	return sum
}

func even(x int) bool {
	return x&1 == 0
}