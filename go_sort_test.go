package gobenchmark

import (
	"math/rand"
	"testing"
)

func genArrInt(n int) []int {
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(n))
	}
	return arr
}

func genArrStruct(n int) []MyStruct {
	arr := []MyStruct{}
	for i := 0; i < n; i++ {
		arr = append(arr, MyStruct{Val: rand.Intn(n)})
	}
	return arr
}

func genArrPtr(n int) []*MyStruct {
	arr := []*MyStruct{}
	for i := 0; i < n; i++ {
		arr = append(arr, &MyStruct{Val: rand.Intn(n)})
	}
	return arr
}

func BenchmarkSliceStable1KInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(1000)
		b.StartTimer()
		SliceStableInt(arr)
	}
}

func BenchmarkSliceStable10KInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(10_000)
		b.StartTimer()
		SliceStableInt(arr)
	}
}

func BenchmarkSliceStable1KStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(1000)
		b.StartTimer()
		SliceStableStruct(arr)
	}
}

func BenchmarkSliceStable10KStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(10_000)
		b.StartTimer()
		SliceStableStruct(arr)
	}
}

func BenchmarkSliceStable1KPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(1000)
		b.StartTimer()
		SliceStablePtr(arr)
	}
}

func BenchmarkSliceStable10KPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(10_000)
		b.StartTimer()
		SliceStablePtr(arr)
	}
}

func BenchmarkSortStable1KInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(1000)
		b.StartTimer()
		SortStableInt(arr)
	}
}

func BenchmarkSortStable10KInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(10_000)
		b.StartTimer()
		SortStableInt(arr)
	}
}

func BenchmarkSortStable1KStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(1000)
		b.StartTimer()
		SortStableStruct(arr)
	}
}

func BenchmarkSortStable10KStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(10_000)
		b.StartTimer()
		SortStableStruct(arr)
	}
}

func BenchmarkSortStable1KPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(1000)
		b.StartTimer()
		SortStablePtr(arr)
	}
}

func BenchmarkSortStable10KPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(10_000)
		b.StartTimer()
		SortStablePtr(arr)
	}
}
