default: build

build:
	go fmt
	go vet
	go build --ldflags="-X github.com/enbritely/heartbeat-golang.CommitHash `git rev-parse HEAD`"

test: build
	go test 

coverage-test:
	go test -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out
