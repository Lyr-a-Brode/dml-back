PKG := "github.com/Lyr-a-Brode/dml-back"

export GOPROXY=direct
export GOSUMDB=off

.PHONY: dep build


all: build

lint:
	golangci-lint run

dep:
	go mod vendor
	go mod verify

build:
	go build -a -o dml-back -v $(PKG)/cmd/dml