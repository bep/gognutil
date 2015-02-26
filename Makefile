
GOFLAGS ?= $(GOFLAGS:)

all: build

build: test
	@go build $(GOFLAGS) -o bin/gognutil main.go 

test: 
	@go test $(GOFLAGS) ./...

clean:
	@go clean $(GOFLAGS) -i ./...

## EOF

