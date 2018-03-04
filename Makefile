default: build

build:
	go build -v -o ./bin/wps ./cmd

install:
	go build -v -o ${GOPATH}/bin/wps ./cmd
