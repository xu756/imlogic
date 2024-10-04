FROM alpine:3.20.3
WORKDIR /app
RUN apk update --no-cache && apk add --no-cache tzdata
ADD main .
ENV TZ Asia/Shanghai

CMD ["./main"]