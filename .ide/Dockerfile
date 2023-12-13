FROM golang:1.21
RUN curl -fsSL https://code-server.dev/install.sh | sh &&\
    code-server --install-extension redhat.vscode-yaml &&\
    code-server --install-extension MS-CEINTL.vscode-language-pack-zh-hans &&\
    code-server --install-extension ms-azuretools.vscode-docker &&\
    code-server --install-extension mtxr.sqltools &&\
    code-server --install-extension mtxr.sqltools-driver-pg &&\
    echo done \

RUN go install github.com/cortesi/modd/cmd/modd@latest

ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8