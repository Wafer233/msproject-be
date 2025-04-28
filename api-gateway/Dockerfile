FROM ubuntu:latest

LABEL authors="wafer"

WORKDIR /app

# 直接复制当前目录下的文件，而不是使用完整路径
COPY cmd/bin/api-gateway /app/api-gateway
COPY config/config.yaml /app/config/config.yaml

# 开放端口
EXPOSE 80

ENTRYPOINT ["/app/api-gateway"]