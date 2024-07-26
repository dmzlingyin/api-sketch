export PROFILE = dev

.PHONY: proto

all:
	go build -o api;./api

proto:
	cp -r ~/project/self/service-sketch/proto .
	cd proto; sh build.sh

tidy:
	go mod tidy