GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: test

lint:
	@if [ ! -f bin/golangci-lint ]; then \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.21.0; \
	fi
	./bin/golangci-lint run ./...

fmt:
	gofmt -w $(GOFMT_FILES)

test: fmt lint
	@sh -c "go test ./... -timeout=2m -parallel=4"

.PHONY: default test cover fmt lint
