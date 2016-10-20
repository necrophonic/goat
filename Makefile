BIN=goat
TESTS?=./...

deps:
	-go get github.com/ogier/pflag
	-go get github.com/necrophonic/log

test-deps:
	-go get github.com/smartystreets/goconvey

build: deps

test-unit: build test-deps
	@go test $(TESTS)
