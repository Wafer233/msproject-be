FROM ubuntu:latest

LABEL authors="wafer"

WORKDIR /app

# 直接复制当前目录下的文件
COPY cmd/bin/project-service /app/project-service
COPY config/config.yaml /app/config/config.yaml

# 开放端口
EXPOSE 8081

ENTRYPOINT ["/app/project-service"]