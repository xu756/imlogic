FROM golang:alpine3.19  as builder
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

ENV GOPATH /go
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w"  -o /app/main  cmd/im/rpc/main.go
RUN ls /app

FROM alpine:3.19.1

WORKDIR /app

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
COPY --from=builder /app/main /app/main
ENV TZ Asia/Shanghai

CMD ["./main", "-f", "deploy.yaml"]