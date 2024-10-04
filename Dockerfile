FROM ubuntu:24.10

WORKDIR /app
RUN apt update --no-cache && apk add --no-cache tzdata
COPY main /app/
ENV TZ Asia/Shanghai

CMD ["app/main"]