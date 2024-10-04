FROM alpine:3.20.3
WORKDIR /app
RUN apk update --no-cache && apk add --no-cache tzdata
COPY main /app
ENV TZ Asia/Shanghai

CMD ["./main"]