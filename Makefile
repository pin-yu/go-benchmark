.PHONY: bench
bench:
	go test -v -bench=. -run=none . -benchmem -cpu 1
