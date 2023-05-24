.PHONY: local server clean lint fmt

APP_NAME=resumme-builder

fmt:
	gofmt -s -l -w .

lint: fmt
	golangci-lint run

local:
	go run *.go local -f='$(file)'

server:
	go run *.go server

build:
	go build -o bin/$(APP_NAME)

docker-build:
	docker build --no-cache -t $(APP_NAME) .

docker-run:
	docker run -it --rm -p 9000:9000 $(APP_NAME)

clean:
	rm -rf output
	rm -rf bin/
	@echo "cleaned!"
