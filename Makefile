
build: fmt
	go build

fmt:
	gofmt -w -s *.go */*.go
