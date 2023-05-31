.PHONY: build

GIT_SHA := $(shell git rev-parse HEAD)

build:
	docker build .		\
	-t tt:latest

build-with-hash:
	docker build . 		\
	-t tt:$(GIT_SHA) 	\
	-t tt:latest

run:
	docker run -it		\
	--rm				\
	--name TT			\
	tt:latest
