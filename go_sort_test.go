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

// non stable parts

func BenchmarkSliceNonStable10Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(10)
		b.StartTimer()
		SliceNonStableInt(arr)
	}
}

func BenchmarkSliceNonStable100Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(100)
		b.StartTimer()
		SliceNonStableInt(arr)
	}
}

func BenchmarkSliceNonStable1KInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(1000)
		b.StartTimer()
		SliceNonStableInt(arr)
	}
}

func BenchmarkSliceNonStable10KInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(10_000)
		b.StartTimer()
		SliceNonStableInt(arr)
	}
}

func BenchmarkSliceNonStable10Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(10)
		b.StartTimer()
		SliceNonStableStruct(arr)
	}
}

func BenchmarkSliceNonStable100Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(100)
		b.StartTimer()
		SliceNonStableStruct(arr)
	}
}

func BenchmarkSliceNonStable1KStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(1000)
		b.StartTimer()
		SliceNonStableStruct(arr)
	}
}

func BenchmarkSliceNonStable10KStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(10_000)
		b.StartTimer()
		SliceNonStableStruct(arr)
	}
}

func BenchmarkSliceNonStable10Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(10)
		b.StartTimer()
		SliceNonStablePtr(arr)
	}
}

func BenchmarkSliceNonStable100Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(100)
		b.StartTimer()
		SliceNonStablePtr(arr)
	}
}

func BenchmarkSliceNonStable1KPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(1000)
		b.StartTimer()
		SliceNonStablePtr(arr)
	}
}

func BenchmarkSliceNonStable10KPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(10_000)
		b.StartTimer()
		SliceNonStablePtr(arr)
	}
}

func BenchmarkSortNonStable10Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(10)
		b.StartTimer()
		SortNonStableInt(arr)
	}
}

func BenchmarkSortNonStable100Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(100)
		b.StartTimer()
		SortNonStableInt(arr)
	}
}

func BenchmarkSortNonStable1KInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(1000)
		b.StartTimer()
		SortNonStableInt(arr)
	}
}

func BenchmarkSortNonStable10KInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(10_000)
		b.StartTimer()
		SortNonStableInt(arr)
	}
}

func BenchmarkSortNonStable10Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(10)
		b.StartTimer()
		SortNonStableStruct(arr)
	}
}

func BenchmarkSortNonStable100Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(100)
		b.StartTimer()
		SortNonStableStruct(arr)
	}
}

func BenchmarkSortNonStable1KStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(1000)
		b.StartTimer()
		SortNonStableStruct(arr)
	}
}

func BenchmarkSortNonStable10KStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(10_000)
		b.StartTimer()
		SortNonStableStruct(arr)
	}
}

func BenchmarkSortNonStable10Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(10)
		b.StartTimer()
		SortNonStablePtr(arr)
	}
}

func BenchmarkSortNonStable100Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(100)
		b.StartTimer()
		SortNonStablePtr(arr)
	}
}

func BenchmarkSortNonStable1KPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(1000)
		b.StartTimer()
		SortNonStablePtr(arr)
	}
}

func BenchmarkSortNonStable10KPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(10_000)
		b.StartTimer()
		SortNonStablePtr(arr)
	}
}

// stable parts

func BenchmarkSliceStable10Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(10)
		b.StartTimer()
		SliceStableInt(arr)
	}
}

func BenchmarkSliceStable100Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(100)
		b.StartTimer()
		SliceStableInt(arr)
	}
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

func BenchmarkSliceStable10Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(10)
		b.StartTimer()
		SliceStableStruct(arr)
	}
}

func BenchmarkSliceStable100Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(100)
		b.StartTimer()
		SliceStableStruct(arr)
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

func BenchmarkSliceStable10Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(10)
		b.StartTimer()
		SliceStablePtr(arr)
	}
}

func BenchmarkSliceStable100Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(100)
		b.StartTimer()
		SliceStablePtr(arr)
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

func BenchmarkSortStable10Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(10)
		b.StartTimer()
		SortStableInt(arr)
	}
}

func BenchmarkSortStable100Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrInt(100)
		b.StartTimer()
		SortStableInt(arr)
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

func BenchmarkSortStable10Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(10)
		b.StartTimer()
		SortStableStruct(arr)
	}
}

func BenchmarkSortStable100Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrStruct(100)
		b.StartTimer()
		SortStableStruct(arr)
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

func BenchmarkSortStable10Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(10)
		b.StartTimer()
		SortStablePtr(arr)
	}
}

func BenchmarkSortStable100Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := genArrPtr(100)
		b.StartTimer()
		SortStablePtr(arr)
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
