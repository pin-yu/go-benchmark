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

It is better to use `string builder`,especially in the for loop, to concatenate the strings. If we write `string + string` in the for loop, the process will takes a long time to  allocate memories for the intermediate string.

## String to []Rune, []Byte

| Items  | Iterations | ns/op | B/op | allocs/op |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| BenchmarkStringNoConversion | 404204 | 297.3 | 0 | 0 |
| BenchmarkStringToRune | 2252 | 51516 | 144000 | 1000 |
| BenchmarkRuneToString | 708 | 166583 | 48000 | 1000 |
| BenchmarkStringToByte | 7906 | 15527 | 48000 | 1000 |
| BenchmarkByteToString | 8918 | 13802 | 48000 | 1000 |

summary

The ns/op of RuneToString is 3 times greater than StringToRune. The ns/op of ByteToString is just slightly slower than that of StringToByte. Just be careful when using string conversion, especially double conversion (string -> rune -> string).
