

GIT_SHA=$(shell git rev-parse HEAD)


docker:
	docker build . \
	-t tt:$(GIT_SHA)