package main

import "testing"

func BenchmarkEuler1BruteForce(b *testing.B) {
	BruteForce(100000)
}
func BenchmarkEuler1Gauss(b *testing.B) {
	GaussSummation(100000)
}