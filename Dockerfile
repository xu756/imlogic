FROM alpine:3.20.3
WORKDIR /app
RUN apk update --no-cache && apk add --no-cache tzdata
RUN chmod +x app/main
COPY main /app/
ENV TZ Asia/Shanghai

CMD ["app/main"]