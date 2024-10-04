FROM alpine:3.20.2

WORKDIR /app
RUN apk update --no-cache && apk add --no-cache tzdata
COPY /usr/local/bin/runc /usr/local/sbin/
COPY /usr/local/bin/runc /usr/bin/
COPY main /app/
ENV TZ Asia/Shanghai

CMD ["app/main"]