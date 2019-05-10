GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)
TEST_RESULTS=/tmp/test-results

default: format build test vet

format:
	go fmt

vet:
	go vet ./...

test:
	mkdir -p ${TEST_RESULTS}
	@go test -coverprofile=${TEST_RESULTS}/unittest.out -v $(GOPACKAGES)
	@go tool cover -html=${TEST_RESULTS}/unittest.out -o ${TEST_RESULTS}/unittest-coverage.html
	rm -f ${TEST_RESULTS}/unittest.out

build:
	GO111MODULE=on go build -o timesum

release:
	mkdir -p release
	GO111MODULE=on go build -o release/timesum-darwin-amd64

run: build
	./dist/timesum-darwin-amd64
