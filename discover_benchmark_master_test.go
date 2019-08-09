package main

import (
	"fmt"
	"testing"

	pb "github.com/brotherlogic/discovery/proto"
)

func benchmarkMasterRegister(i, d int, b *testing.B) {
	s := InitTestServer()

	testdata := []*pb.RegistryEntry{}

	for c := 0; c < i; c++ {
		testdata = append(testdata, &pb.RegistryEntry{Name: fmt.Sprintf("Server-%v", c), Identifier: fmt.Sprintf("Identifier-%v", c/d), TimeToClean: 100, Master: true})
	}

	for n := 0; n < b.N; n++ {
		run(s, testdata[n%len(testdata)])
	}
}

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
