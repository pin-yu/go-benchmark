.PHONY: bench
bench:
	go test -v -bench=. -run=none . -benchmem -cpu 1 -benchtime=0.1s
