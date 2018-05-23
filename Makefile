

all:
	go fmt 1.go
	GOARCH=arm go build 1.go
