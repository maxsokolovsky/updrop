.DEFAULT_GOAL := build

build:
	go build

build-pi:
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-s -w"
