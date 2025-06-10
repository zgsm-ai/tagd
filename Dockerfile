FROM golang:1.24.0 AS builder
WORKDIR /app
COPY . .

RUN go env -w CGO_ENABLED=0 && \
    go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,direct  

#
#go env -w GOPROXY=http://mirrors.sangfor.org/nexus/repository/go-proxy-group
#

RUN go mod tidy && go build -o tagd *.go

FROM centos:7.6.1810
#时区设置
ENV env prod
ENV TZ Asia/Shanghai
WORKDIR /
COPY --from=builder /app/tagd /usr/local/bin
RUN chmod 755 /usr/local/bin/tagd
ENTRYPOINT ["/usr/local/bin/tagd"]
