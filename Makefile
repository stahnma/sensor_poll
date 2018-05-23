
all: sensor_poll

format:
	go fmt *.go

sensor_poll: format clean
	GOARCH=arm go build sensor_poll.go

install:
	mkdir -p $(DESTDIR)/usr/local/bin
	install -m755 sensor_poll $(DESTDIR)/usr/local/bin
	mkdir -p $(DESTDIR)/etc/systemd/system
	install -m0644 ext/sensor_poll.service $(DESTDIR)/etc/systemd/system/sensor_poll.service
	systemctl daemon-reload
	systemctl enable sensor_poll
	systemctl restart sensor_poll

deps:
	go get github.com/morus12/dht22

uninstall:
	systemctl stop sensor_poll || true
	systemctl disable sensor_poll || true
	rm -f $(DESTDIR)/usr/local/bin/sensor_poll $(DESTDIR)/etc/systemd/system/sensor_poll.service || true
	systemctl daemon-reload || true
clean:
	rm -f sensor_poll
