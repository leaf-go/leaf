FROM golang:1.17-alpine as builder
#  docker build --build-arg env=debug --rm -t leaf-app .

# 工作目录 需要与执行阶段目录一致、为了方便测试
WORKDIR /app

# 添加执行用户 自行添加
RUN adduser -u 10001 -D leaf

# 设置ENV
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

# 这里更新依赖
COPY go.mod .
COPY go.sum .

RUN go mod download

# 将目录复制到容器 /build
COPY . .

# 构建
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o leaf-app .

# 执行阶段
FROM alpine:3.15 as final

# 默认使用gin框架 如果是iris 请自行实现x.IApplication 接口并在boot.services文件中中挂载。
#debug|test|release
ARG env=debug

ENV GIN_MODE $env

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=builder /app/leaf-app /app/
COPY --from=builder /app/config /app/config
COPY --from=builder /etc/passwd /etc/passwd

USER app-runner

EXPOSE 8300

ENTRYPOINT ["/app/leaf-app"]






