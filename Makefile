

all: clean
	go fmt *.go
	GOARCH=arm go build sensor_poll.go

clean:
	rm -f sensor_poll
