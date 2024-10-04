FROM alpine:3.20.3
WORKDIR /app
RUN apk update --no-cache && apk add --no-cache tzdata
COPY deploy/* /app/
RUN chmod +x /app/main
RUN ls -al /app
ENV TZ Asia/Shanghai

CMD ["./main"]