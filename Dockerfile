FROM golang:1.19rc2

ARG TZ
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

COPY ./ /app

RUN go mod download

ENTRYPOINT go run main.go
