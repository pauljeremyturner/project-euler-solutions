package main

import "testing"

func BenchmarkEuler1BruteForce(b *testing.B) {
	for i := 0; i < 1000; i++ {
		BruteForce(100000)
	}
}
func BenchmarkEuler1Gauss(b *testing.B) {
	for i := 0; i < 1000; i++ {
		GaussSummation(100000)
	}
}