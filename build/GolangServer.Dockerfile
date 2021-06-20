FROM golang:latest AS build

WORKDIR /go/src/testTask

COPY . .

RUN GOPATH=/go CGO_ENABLED=0  GOOS=linux  go build -o /main   ./cmd/main

FROM ubuntu:18.04 AS release
MAINTAINER Alex Sirmais

WORKDIR /app
COPY --from=build /main .
RUN chmod +x ./main

EXPOSE 8080/tcp

USER root
CMD sleep 2 && ./main