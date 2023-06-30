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
	-p 10000:10000		\
	--rm				\
	--name TT			\
	tt:latest
