.PHONY: local server clean lint fmt

fmt:
	gofmt -s -l -w .

lint: fmt
	golangci-lint run

local:
	go run *.go local

server:
	go run *.go server


clean:
	rm -rf output
	@echo "Output directory cleaned"
