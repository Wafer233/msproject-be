FROM ubuntu:latest

LABEL authors="wafer"

WORKDIR /app

# 直接复制当前目录下的文件
COPY cmd/bin/user-service /app/user-service
COPY config/config.yaml /app/config/config.yaml

# 开放端口
EXPOSE 8080

ENTRYPOINT ["/app/user-service"]