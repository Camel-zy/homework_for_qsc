FROM golang:1.15 AS builder

ARG CI_JOB_TOKEN=$CI_JOB_TOKEN

ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct \
    GO111MODULE=on \
    WORKDIR=/tmp/workdir/ \
    CGO_ENABLED=0 \
    GOPRIVATE=git.zjuqsc.com/rop/rop-sms

RUN mkdir -p $WORKDIR

RUN git config --global url."https://gitlab-ci-token:${CI_JOB_TOKEN}@git.zjuqsc.com/".insteadOf "https://git.zjuqsc.com/"

COPY go.mod go.sum $WORKDIR

RUN cd $WORKDIR && go mod download all

COPY . $WORKDIR

RUN cd $WORKDIR && go build -o /ropd

FROM alpine:3.13.1

COPY --from=builder /ropd /ropd

CMD ["/ropd"]
