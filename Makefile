.PHONY: iso
iso:
	go run cmd/main.go

test: lint shellcheck

lint:
	golint ./...

shellcheck:
	shellcheck $(shell find . -name '*.sh')
