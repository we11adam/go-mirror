.PHONY: build push

all: build push

build:
	docker build -t ghcr.io/we11adam/go-mirror:latest .

push:
	docker push ghcr.io/we11adam/go-mirror
