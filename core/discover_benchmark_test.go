package discovery

import (
	"context"
	"fmt"
	"log"
	"testing"

	pb "github.com/brotherlogic/discovery/proto"
)

func benchmarkRegister(i, d int, b *testing.B) {
	s := InitTestServer()

	testdata := []*pb.RegistryEntry{}

	for c := 0; c < i; c++ {
		testdata = append(testdata, &pb.RegistryEntry{Name: fmt.Sprintf("Server-%v", c), Identifier: fmt.Sprintf("Identifier-%v", c/d), TimeToClean: 100})
	}

	for n := 0; n < b.N; n++ {
		resp, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: testdata[n%len(testdata)]})
		if err != nil || resp.Service.Port < 50056 || resp.Service.Port > 65535 {
			log.Fatalf("Bad port %v %v", err, resp)
		}
	}
}

func BenchmarkRegister1_5(b *testing.B) {
	benchmarkRegister(1, 5, b)
}

func BenchmarkRegister10_5(b *testing.B) {
	benchmarkRegister(10, 5, b)
}

func BenchmarkRegister100_5(b *testing.B) {
	benchmarkRegister(100, 5, b)
}

func BenchmarkRegister1_10(b *testing.B) {
	benchmarkRegister(1, 10, b)
}

func BenchmarkRegister10_10(b *testing.B) {
	benchmarkRegister(10, 10, b)
}

func BenchmarkRegister100_10(b *testing.B) {
	benchmarkRegister(100, 10, b)
}
