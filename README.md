# Go Benchmark

Benchmark common patterns in Go to figure out what has better performance

- Environment
  - Macbook Air M2 chip with 8GB RAM and 256GB SSD
  - benchmark for 0.1 second

## Build String

| Items  | Iterations | ns/op | B/op | allocs/op |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| BenchmarkStringAdd | 1962 | 57006 | 530274 | 999 |
| BenchmarkStringBuilder | 51897 | 2291 | 3320 | 9 |

summary

It is better to use `string builder`, especially in the for loop, to concatenate the strings. If we write `string + string` in the for loop, the process will takes a long time to  allocate memories for the intermediate strings.

## Convert string to []Rune, []Byte

| Items  | Iterations | ns/op | B/op | allocs/op |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| BenchmarkStringNoConversion | 404204 | 297.3 | 0 | 0 |
| BenchmarkStringToRune | 2252 | 51516 | 144000 | 1000 |
| BenchmarkRuneToString | 708 | 166583 | 48000 | 1000 |
| BenchmarkStringToByte | 7906 | 15527 | 48000 | 1000 |
| BenchmarkByteToString | 8918 | 13802 | 48000 | 1000 |

summary

The ns/op of RuneToString is 3 times greater than StringToRune. The ns/op of ByteToString is just slightly slower than that of StringToByte. Just be careful when using string conversion, especially double conversion (string -> rune -> string).

## Multiplication and Division

| Items  | Iterations | ns/op | B/op | allocs/op |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| BenchmarkMultiplyFloat64 | 3859906 | 30.46 | 0 | 0 |
| BenchmarkDivideFloat64 | 3769795 | 30.97 | 0 | 0 |
| BenchmarkMultiplyInt64 | 3946886 | 30.36 | 0 | 0 |
| BenchmarkDivideInt64 | 3856773 | 31.01 | 0 | 0 |

summary

There is no significant difference among each benchmarks.

## Slice

| Items  | Iterations | ns/op | B/op | allocs/op |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| BenchmarkSetValueToSliceOf0Len0Cap        | 49  | 2247754 | 41678080 | 38 |
| BenchmarkSetValueToSliceOfNLenNCap        | 248 | 487447 | 8003584 | 1 |
| BenchmarkSetValueToSliceOf0LenNCap        | 196 | 593774 | 8003584 | 1 |
| BenchmarkSetPtrToSliceOf0Len0Cap          | 4 | 32370386 | 49678092 | 1000038 |
| BenchmarkSetPtrToSliceOfNLenNCap          | 8 | 15173344 | 16003590 | 1000001 |
| BenchmarkSetPtrToSliceOf0LenNCap          | 9 | 14534449 | 16003587 | 1000001 |
| BenchmarkSetPtrToInterfaceSliceOf0Len0Cap | 2 | 53907958 | 96036600 | 1000039 |
| BenchmarkSetPtrToInterfaceSliceOfNLenNCap | 8 | 15025208 | 24007170 | 1000001 |
| BenchmarkSetPtrToInterfaceSliceOf0LenNCap | 9 | 13565810 | 24007173 | 1000001 |

summary

1. The access speed is the fastest if we know the expected length of a slice and use `make([]struct{}, n)` to initialize length and capacity at first.
2. Be aware of using slices of pointer, because they are extremely slow.
3. Using slices of interface is the worst case.

## Structure padding

Go will pad the structure to their largest field alignment guarantees. See `T5` in `structure_padding.go` for more information.
