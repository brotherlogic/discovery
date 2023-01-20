package main

import (
	"testing"
)

func BenchmarkMasterRegister1_5(b *testing.B) {
	benchmarkRegister(1, 5, b)
}

func BenchmarkMasterRegister10_5(b *testing.B) {
	benchmarkRegister(10, 5, b)
}

func BenchmarkMasterRegister100_5(b *testing.B) {
	benchmarkRegister(100, 5, b)
}

func BenchmarkMasterRegister1000_5(b *testing.B) {
	benchmarkRegister(1000, 5, b)
}

func BenchmarkMasterRegister5000_5(b *testing.B) {
	benchmarkRegister(1000, 5, b)
}

func BenchmarkMasterRegister1_10(b *testing.B) {
	benchmarkRegister(1, 10, b)
}

func BenchmarkMasterRegister10_10(b *testing.B) {
	benchmarkRegister(10, 10, b)
}

func BenchmarkMasterRegister100_10(b *testing.B) {
	benchmarkRegister(100, 10, b)
}
func BenchmarkMasterRegister1000_10(b *testing.B) {
	benchmarkRegister(1000, 10, b)
}

func BenchmarkMasterRegister5000_10(b *testing.B) {
	benchmarkRegister(1000, 10, b)
}
