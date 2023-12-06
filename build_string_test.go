package gobenchmark

import (
	"testing"
)

func BenchmarkStringAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringAdd()
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuilder()
	}
}
