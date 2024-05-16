package gobenchmark

func MultiplyFloat64() float64 {
	var n float64 = 1e300
	for i := 0; i < 100; i++ {
		n *= 0.1
	}
	return n
}

func DivideFloat64() float64 {
	var n float64 = 1e300
	for i := 0; i < 100; i++ {
		n /= 10
	}
	return n
}

func MultiplyInt64() int64 {
	var n int64 = 1
	for i := 0; i < 100; i++ {
		n *= 3
	}
	return n
}

func DivideInt64() int64 {
	var n int64 = 3 ^ 101
	for i := 0; i < 100; i++ {
		n /= 3
	}
	return n
}
