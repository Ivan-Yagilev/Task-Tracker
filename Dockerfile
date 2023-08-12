FROM golang:alpine

WORKDIR /app

COPY ./ ./

RUN go build -o main cmd/main.go

ENTRYPOINT ./main