default: build

build: 
	go build --ldflags="-X github.com/enbritely/heartbeat-golang.CommitHash `git rev-parse HEAD`"
