FROM golang:1.15 AS builder

ARG REPO_USER=$REPO_USER
ARG REPO_PASSWD=$REPO_PASSWD

ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct \
    GO111MODULE=on \
    WORKDIR=/tmp/workdir/ \
    CGO_ENABLED=0 \
    GOPRIVATE=git.zjuqsc.com/rop/rop-sms

RUN mkdir -p $WORKDIR

RUN git config --global url."https://${REPO_USER}:${REPO_PASSWD}@git.zjuqsc.com/".insteadOf "https://git.zjuqsc.com/"

COPY go.mod go.sum $WORKDIR

RUN cd $WORKDIR && go mod download all

COPY . $WORKDIR

RUN cd $WORKDIR && go build -o /ropd

FROM alpine:3.13.1

COPY --from=builder /ropd /ropd

CMD ["/ropd"]
