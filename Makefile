PKG_NAME=gans
DUSER=aceberg

mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/gans && \
	go mod tidy

run:
	cd cmd/gans/ && \
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint

go-build:
	cd cmd/gans/ && \
	CGO_ENABLED=0 go build -o ../../tmp/gans .

docker-build:
	docker build -t $(DUSER)/$(PKG_NAME) .

cli-run:
	cd cmd/gans-cli/ && \
	go run .