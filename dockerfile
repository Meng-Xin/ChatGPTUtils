# 编译阶段：引用最小编译环境
FROM golang:alpine AS builder

# 镜像默认工作目录
WORKDIR /root/code

# 配置镜像golang的默认配置,方便拉取依赖
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 拷贝当前目录所有文件到工作目录
COPY . .

# 设置编译环境并进行编译
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64  go build -o app .

# 构建阶段：使用 alpine 最小构建
FROM alpine:latest

# 设置镜像工作目录
WORKDIR /go/gocode

# 在builder阶段复制可执行的go二进制文件app
COPY --from=builder ./root/code .

# 时区设置
ENV TZ="Asia/Shanghai"

# 开放端口 80
EXPOSE 80

# 启动服务
ENTRYPOINT ["./app"]
# 或者启动
CMD ["app"]