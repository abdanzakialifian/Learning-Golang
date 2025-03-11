package test

import (
	"learning-benchmark-unit-test-golang/helper"
	"testing"
)

func BenchmarkSayHello(b *testing.B) {
	b.Run("Benchmark1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helper.SayHello("Abdan")
		}
	})

	b.Run("Benchmark2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helper.SayHello("Zaki")
		}
	})

	b.Run("Benchmark3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helper.SayHello("Alifian")
		}
	})

	b.Run("Benchmark4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helper.SayHello("Golang")
		}
	})

	b.Run("Benchmark5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helper.SayHello("Benchmark")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Benchmark1",
			request: "Abdan",
		},
		{
			name:    "Benchmark2",
			request: "Zaki",
		},
		{
			name:    "Benchmark3",
			request: "Alifian",
		},
		{
			name:    "Benchmark4",
			request: "Golang",
		},
		{
			name:    "Benchmark5",
			request: "Benchmark",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				helper.SayHello(benchmark.request)
			}
		})
	}
}
