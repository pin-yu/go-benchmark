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

### Summary

It is better to use `string builder`,especially in the for loop, to concatenate the strings. If we write `string + string` in the for loop, the process will takes a long time to  allocate memories for the intermediate string.
