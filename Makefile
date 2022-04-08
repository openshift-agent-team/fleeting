NODE_ZERO_IP ?= 192.168.122.2

.PHONY: iso
iso: clean
	go run cmd/main.go

.PHONY: clean realclean
clean:
	rm -rf bin
	rm -f output/fleeting.iso

realclean: clean
	rm -rf output

.PHONY: test
test: lint shellcheck

.PHONY: lint
lint:
	golint ./...

.PHONY: shellcheck
shellcheck:
	shellcheck $(shell find . -name '*.sh')
