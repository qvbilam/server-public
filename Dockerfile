FROM golang:1.20.14-alpine AS go-alpine-builder
# 拷贝代码
COPY . /go/src/public
# 设置工作目录
WORKDIR /go/src/public
# 设置执行命令
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go build -o app

# 重制作镜像-通过dockerhub可查看golang1.18-alpine对应的alpine版本
FROM alpine:3.15
COPY --from=go-alpine-builder /go/src/public/app /bin/public
ENV PORT=9501

# 暴露端口
EXPOSE 9501

ENTRYPOINT [ "/bin/public" ]
