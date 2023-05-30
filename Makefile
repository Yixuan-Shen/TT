.PHONY: build

GIT_SHA=$(shell git rev-parse HEAD)


build:
	docker build . \
	-t tt:$(GIT_SHA) \
	-t tt:latest