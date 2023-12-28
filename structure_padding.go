package gobenchmark

// Architecture: MacOS with M2 cores

// total 1 bytes
type T1 struct {
	A int8 // 1 bytes
}

// total 2 bytes
type T2 struct {
	A int8 // 1 bytes
	B int8 // 1 bytes
}

// total 3 bytes
type T3 struct {
	A int8 // 1 bytes
	B int8 // 1 bytes
	C int8 // 1 bytes
}

// total 16 bytes
type T4 struct {
	A int8  // 1 bytes + 7 padding bytes
	B int64 // 8 bytes
}

// The largest field alignment guarantee of T5 is 8 bytes (int64).
// Therefore, T5 will be padded as the multiple of 8, which is 24 bytes.
// total 24 bytes
type T5 struct {
	A int8  // 1 bytes + 7 padding bytes
	B int64 // 8 bytes
	C int16 // 2 bytes + 6 padding bytes
}
