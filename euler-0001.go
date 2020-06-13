package main

import (
	"fmt"
	"math"
)

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

func BruteForce(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		if (0 == i%3) || (0 == i%5) {
			count += i
		}
	}
	return count
}


func GaussSummation(x int) int {
	n  := float64(x)
	fizz := math.Floor(n / 3)
	buzz := math.Floor(n / 5)
	fizzbuzz := math.Floor(n / 15)

	result := 0.5 * ( 3 * fizz * (fizz + 1) + 5 * buzz * (buzz + 1) - 15 * fizzbuzz * (fizzbuzz + 1))
	return int(result)

}


func main() {

	fmt.Println(BruteForce(1000))
	fmt.Println(GaussSummation(1000))

}