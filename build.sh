#!/bin/bash

# 检查环境变量DOCKER_IMAGE_NAME是否已设置
if [ -z "${DOCKER_IMAGE_NAME}" ]; then
  echo "DOCKER_IMAGE_NAME环境变量未设置。"
  exit 1
fi

# 检查Dockerfile是否存在
if [ ! -f "Dockerfile" ]; then
  echo "Dockerfile不存在。"
  exit 1
fi
#
## 显示旧Dockerfile内容及颜色
#echo -e "\033[32m 旧Dockerfile内容如下： \033[0m"
#echo "---------------------------------------------------"
#cat Dockerfile
#echo "---------------------------------------------------"
#
## 根据DOCKER_IMAGE_NAME设置mapped_value
#case "${DOCKER_IMAGE_NAME}" in
#  "user-rpc") mapped_value="user" ;;
#  "api") mapped_value="api" ;;
#  "im-rpc") mapped_value="im/rpc" ;;
#  "im-server") mapped_value="im/server" ;;
#  *)
#    echo "没有找到${DOCKER_IMAGE_NAME}的映射值。"
#    exit 2
#    ;;
#esac
#
## 获取Dockerfile的完整路径
#dockerfile_path="$(pwd)/Dockerfile"
#
## 删除Dockerfile中的第10行并插入新内容
#if sed -i  '15d' "${dockerfile_path}" && \
#  sed -i  "14a RUN go build -o main -v cmd/${mapped_value}/main.go" "${dockerfile_path}"; then
#  echo "---------------------------------------------------"
#  echo -e "\033[32m Dockerfile已更新。 \033[0m"
#  cat Dockerfile
#  echo "Dockerfile已更新。"
#else
#  echo "更新Dockerfile时发生错误。"
#fi
