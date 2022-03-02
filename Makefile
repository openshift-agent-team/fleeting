.PHONY: iso

clean:
	rm -rf bin
	rm -rf output

iso: clean
	go run cmd/main.go -node-zero-ip 192.168.122.2

test: lint shellcheck

lint:
	golint ./...

shellcheck:
	shellcheck $(shell find . -name '*.sh')
