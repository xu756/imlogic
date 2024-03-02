FROM golang:alpine3.19  as builder

RUN mkdir /app

ADD configs/dev.yaml /app/dev.yaml
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
ENV GOPATH=/go
ENV GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"
RUN go mod tidy
RUN go build -o main -v cmd/im/rpc/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/dev.yaml .

CMD ["/app/main -f /app/deploy.yaml"]