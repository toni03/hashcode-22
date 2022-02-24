# Makefile for hashcode-21
# License MIT

# Variables
UNAME		:= $(uname -s)
APP_NAME	:= hashcode
APP_VERSION	:= 0.0.1
GO_FLAGS	:= CGO_ENABLED=0
LD_FLAGS	:= --ldflags '-X main.version=${APP_VERSION} -X main.appName=${APP_NAME} -extldflags "-static" -w -s'

.PHONY: help
help:	### This screen. Keep it first target to be default
ifeq ($(UNAME), Linux)
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
else
	@# this is not tested, but prepared in advance for you, Mac drivers
	@awk -F ':.*###' '$$0 ~ FS {printf "%15s%s\n", $$1 ":", $$2}' \
		$(MAKEFILE_LIST) | grep -v '@awk' | sort
endif

# Targets
.PHONY: deps deps-sync fmt vet lint clean build install run clean-test test cover bench
deps: ### Get dependencies
	go get -u -v
deps-sync: ### Synchronize dependencies
	go mod vendor
fmt: ### Format code
	gofmt -s -w .
vet: ### Report suspicious code problems
	go vet ./...
lint: fmt vet ### Format and vet code
clean: ### Clean binary
	rm -rf bin/
build: clean ### Build binary
	$(GO_FLAGS) go build -o bin/$(APP_NAME) -a $(LD_FLAGS) -v .
install: ### Install binary in your system
	go install -v .
run: ### Run project for development
	go run . $(ARGS)
clean-test: ### Clean test cache
	go clean -testcache ./...
test: lint ### Run tests
	go test -v ./... -coverprofile=cover.out
cover: test ### Run tests and generate coverage
	go tool cover -html=cover.out -o=cover.html
mocks: ### Generates mocks for interfaces
	mockery --all --keeptree
bench: ### Run tests and benchmarks
	go test -v ./... -benchmem -bench=.