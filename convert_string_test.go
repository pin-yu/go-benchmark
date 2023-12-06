package gobenchmark

import "testing"

func BenchmarkStringNoConversion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringNoConversion()
	}
}

func BenchmarkStringToRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToRune()
	}
}

func BenchmarkRuneToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RuneToString()
	}
}

func BenchmarkStringToByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToByte()
	}
}

func BenchmarkByteToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByteToString()
	}
}
