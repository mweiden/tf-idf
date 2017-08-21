default: build

.PHONY: clean
clean:
	git clean -fd

.PHONY: build
build:
	go build

.PHONY: test
test:
	go test
