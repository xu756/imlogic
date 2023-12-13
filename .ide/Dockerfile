FROM golang:1.21
RUN go install github.com/cortesi/modd/cmd/modd@latest


# 指定字符集支持命令行输入中文（根据需要选择字符集）
ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8