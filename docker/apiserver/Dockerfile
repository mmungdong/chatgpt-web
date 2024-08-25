# base image
FROM node:latest AS frontend

# 安装 pnpm
RUN npm install -g pnpm

# 克隆 GitHub 仓库
RUN git clone -b dev https://github.com/mmungdong/chatgpt-web-dashboard.git /app

# 设置工作目录
WORKDIR /app/chatgpt-web-dashboard

# 安装项目依赖
RUN pnpm install

# 执行构建命令
RUN pnpm run build

FROM golang:1.20.0 AS serverBuilder
ARG VERSION=v0.0.0

WORKDIR /src
COPY ../../../go.mod ../../../go.sum .
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download
COPY ../../../.. /src
RUN mkdir -p bin/ && go build -ldflags "-X main.Version=$VERSION" -o ./bin/ ./app/prom_server/...

FROM alpine:latest

COPY --from=frontend /app/chatgpt-web-dashboard/dist /app/public

# COPY web/admin /app/public/admin

# COPY web/static /app/public/static

COPY ./../dist/server /app/server

ADD dist/server /app/server

RUN mkdir -p /data

EXPOSE 7080

CMD ["/app/server", "--frontend-path=/app/public", "--db=root:123456@tcp(127.0.0.1:3306)/chatgpt_web?charset=utf8mb4&parseTime=True&loc=Local", "--auth-user=admin", "--auth-password=123456"]