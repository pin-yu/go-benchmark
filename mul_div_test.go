package gobenchmark

import "testing"

func BenchmarkMultiplyFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiplyFloat64()
	}
}

func BenchmarkDivideFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideFloat64()
	}
}
func BenchmarkMultiplyInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiplyInt64()
	}
}
func BenchmarkDivideInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideInt64()
	}
}
