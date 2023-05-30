ARG BASE=golang:1.20-alpine3.18
FROM $BASE AS builder

CMD [ "go", "version" ]

