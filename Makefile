.PHONY: build
build:
	cd cmd/watchdog && GOOS=linux GOARCH=amd64 go build -v
.PHONY: run
run:
	cd cmd/watchdog && go build && ./watchdog
.PHONY: clean
clean:
	rm cmd/watchdog/watchdog