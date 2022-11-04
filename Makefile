.PHONY: output clean lint

fmt:
	gofmt -s -l .

lint: fmt
	golangci-lint run

output:
	go run cmd/local/main.go

run:
	go run cmd/server/main.go basic

clean:
	rm -rf output