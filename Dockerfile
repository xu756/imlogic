FROM golang:1.22.0 as builder

RUN mkdir /app

ADD deploy/build /app/

WORKDIR /app
ENV GOPATH=/go
ENV GOPROXY https://mirrors.aliyun.com/goproxy/,direct
RUN go build -o main -v cmd/im/rpc/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["/app/main -f /app/deploy.yaml"]