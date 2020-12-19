FROM golang:1.15

ENV GOPROXY=https://goproxy.io,direct
ENV GO111MODULE=on

RUN mkdir -p /tmp/workdir
COPY . /tmp/workdir

RUN cd /tmp/workdir && go build -o /ropd

CMD ["/ropd"]
