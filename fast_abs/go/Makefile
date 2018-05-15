GO=go

.PHONY: all build run test bench

all: fast_abs

build:
	$(GO) build fast_abs.go

run:
	$(GO) run fast_abs.go

test:
	$(GO) test .

bench:
	$(GO) test . -bench -benchmem
