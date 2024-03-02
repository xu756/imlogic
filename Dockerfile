FROM golang:alpine3.19  as builder

RUN mkdir /app
ADD . /app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
ENV GOPATH=/go
ENV GOPROXY="https://goproxy.cn,direct"

RUN go build -o main  ./cmd/im/rpc/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main

CMD ["./main", "-f", "./deploy.yaml"]