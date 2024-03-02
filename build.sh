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

# 显示旧Dockerfile内容 颜色
echo -e "\033[32m 旧Dockerfile内容如下： \033[0m"
echo  "---------------------------------------------------"
cat Dockerfile

# 根据DOCKER_IMAGE_NAME设置mapped_value
if [ "${DOCKER_IMAGE_NAME}" == "user-rpc" ]; then
  mapped_value="user"
elif [ "${DOCKER_IMAGE_NAME}" == "api" ]; then
  mapped_value="api"
elif [ "${DOCKER_IMAGE_NAME}" == "im-rpc" ]; then
  mapped_value="im/rpc"
elif [ "${DOCKER_IMAGE_NAME}" == "im-server" ]; then
  mapped_value="im/server"
else
  echo "没有找到${DOCKER_IMAGE_NAME}的映射值。"
  exit 2
fi
# 如果没有找到映射值，退出脚本
if [ -z "${mapped_value}" ]; then
  echo "没有找到${DOCKER_IMAGE_NAME}的映射值。"
  exit 2
fi

sed -i '' "10d" Dockerfile
sed -i '' "9a RUN go build -o main -v cmd/${mapped_value}/main.go" Dockerfile

echo  "---------------------------------------------------"
echo -e "\033[32m Dockerfile已更新。 \033[0m"
cat Dockerfile
echo "Dockerfile已更新。"
