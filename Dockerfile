FROM golang:1.19rc2

WORKDIR /app

COPY ./ /app

RUN go mod download

ENTRYPOINT go run main.go
