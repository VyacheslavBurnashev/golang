FROM golang:latest 

RUN mkdir /build
WORKDIR /build
COPY ./ /golang
RUN go mod download

EXPOSE 8081

ENTRYPOINT go run golang/main.go