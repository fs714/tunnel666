.PHONY: build

default: build

BINARY=tunnel666
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-s -X main.BuildTime=${BUILD_TIME}"

build:
	env GOOS=linux GOARCH=amd64 go build -o bin/${BINARY} ${LDFLAGS}
clean:
	rm -rf bin/
