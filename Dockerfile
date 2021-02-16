FROM golang:1.15 AS builder

ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct \
    GO111MODULE=on \
    WORKDIR=/tmp/workdir/ \
    CGO_ENABLED=0

RUN mkdir -p $WORKDIR

COPY go.mod go.sum $WORKDIR

RUN cd $WORKDIR && go mod download all

COPY . $WORKDIR

RUN cd $WORKDIR && go build -o /ropd

FROM alpine:3.13.1

COPY --from=builder /ropd /ropd

CMD ["/ropd"]
