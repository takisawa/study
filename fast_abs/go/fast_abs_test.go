package main

import (
	"testing"
)

const n = 1000000

func BenchmarkFastAbs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := -n; j < 0; j++ {
			fastAbs(j)
		}
	}
}

func BenchmarkSlowAbs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := -n; j < 0; j++ {
			slowAbs(j)
		}
	}
}
