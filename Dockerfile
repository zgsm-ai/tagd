FROM golang:1.24.0 AS builder
WORKDIR /app
COPY . .

RUN go env -w CGO_ENABLED=0 && \
    go env -w GO111MODULE=on && \
    go env -w GOPROXY=http://mirrors.sangfor.org/nexus/repository/go-proxy-group,https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,direct  

#
#go env -w GOPROXY=http://mirrors.sangfor.org/nexus/repository/go-proxy-group
#

RUN go mod tidy 
RUN go build -ldflags="-s -w" -o tagd *.go
RUN chmod 755 tagd

FROM alpine:3.21 AS runtime
ENV env prod
ENV TZ Asia/Shanghai
WORKDIR /
COPY --from=builder /app/tagd /usr/local/bin
ENTRYPOINT ["/usr/local/bin/tagd"]
