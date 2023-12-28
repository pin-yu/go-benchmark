package gobenchmark

import "testing"

func BenchmarkSetValueToSliceOf0Len0Cap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetValueToSliceOf0Len0Cap(1000_000)
	}
}

func BenchmarkSetValueToSliceOfNLenNCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetValueToSliceOfNLenNCap(1000_000)
	}
}

func BenchmarkSetValueToSliceOf0LenNCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetValueToSliceOf0LenNCap(1000_000)
	}
}

func BenchmarkSetPtrToSliceOf0Len0Cap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetPtrToSliceOf0Len0Cap(1000_000)
	}
}

func BenchmarkSetPtrToSliceOfNLenNCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetPtrToSliceOfNLenNCap(1000_000)
	}
}

func BenchmarkSetPtrToSliceOf0LenNCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetPtrToSliceOf0LenNCap(1000_000)
	}
}

func BenchmarkSetPtrToInterfaceSliceOf0Len0Cap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetPtrToInterfaceSliceOf0Len0Cap(1000_000)
	}
}

func BenchmarkSetPtrToInterfaceSliceOfNLenNCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetPtrToInterfaceSliceOfNLenNCap(1000_000)
	}
}

func BenchmarkSetPtrToInterfaceSliceOf0LenNCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetPtrToInterfaceSliceOf0LenNCap(1000_000)
	}
}
