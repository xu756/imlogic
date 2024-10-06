FROM golang:alpine3.20  as builder
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
ENV GOOS=linux
ENV GOARCH=amd64

ENV GOPATH /go
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache tzdata

# 设置工作目录

COPY . /app
WORKDIR /app



RUN go build -ldflags="-s -w"  -o /app/main  app/im/handler/main.go


FROM alpine:3.20.2

WORKDIR /app

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
COPY --from=builder /app/main /app/main
ENV TZ Asia/Shanghai
ENV CONFIG_PATH /config/deploy.yaml

CMD ["./main"]