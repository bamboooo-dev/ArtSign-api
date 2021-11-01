REVISION := $(shell git rev-parse --short HEAD)

default:
	go build -o bin/artsign -ldflags "-X 'main.revision=$(REVISION)'" \
		cmd/artsign/main.go

gen:
	go generate ./...

run:
	go run ./cmd/artsign
