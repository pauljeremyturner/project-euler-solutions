package main

import "testing"

func BenchmarkEuler2Simple(b *testing.B) {
	for i := 0 ; i < 1000; i++ {
		SumEvenFibonacciSimple(1000)
	}
}

func BenchmarkEuler2Bitshift(b *testing.B) {
	for i := 0 ; i < 1000; i++ {
		SumEvenFibonacciBitshift(1000)
	}
}
