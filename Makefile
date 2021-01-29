.PHONY: build
build:
	cd cmd/watchdog && go build -v 
.PHONY: run
run:
	cd cmd/watchdog && go build && ./watchdog
.PHONY: clean
clean:
	rm cmd/watchdog/watchdog