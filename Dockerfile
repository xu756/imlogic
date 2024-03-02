FROM golang:1.20 as builder

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["/app/main"]