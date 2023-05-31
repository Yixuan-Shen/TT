ARG BASE=golang:1.20-alpine3.18
FROM $BASE AS builder

ARG ADD_BUILD_TAGS=""
# ARG MAKE="make -e ADD_BUILD_TAGS=$ADD_BUILD_TAGS build"

ARG ALPINE_PKG_BASE="make git nginx"
ARG ALPINE_PKG_EXTRA=""

LABEL Name=tt-is-more-than-a-traker

RUN apk add --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

WORKDIR /tt-is-more-than-a-traker

# COPY go.mod vendor* ./
RUN [ ! -d "vendor" ] && go mod download all || echo "skipping..."

COPY . .

# RUN ${MAKE}
