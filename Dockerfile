FROM alpine:3.20.2

WORKDIR /app
RUN apk update --no-cache && apk add --no-cache tzdata
COPY deploy/main /app/main
ENV TZ Asia/Shanghai

CMD ["./main"]