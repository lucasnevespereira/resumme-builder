.PHONY: output clean lint

fmt:
	gofmt -s -l .

lint: fmt
	golangci-lint run

output:
	go run .

clean:
	rm -rf output