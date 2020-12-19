FROM golang:1.15

ENV GOPROXY=https://goproxy.io,direct \
    GO111MODULE=on \
    WORKDIR=/tmp/workdir/

RUN mkdir -p $WORKDIR

COPY go.mod go.sum $WORKDIR

RUN go mod download all

COPY . $WORKDIR

RUN cd $WORKDIR && go build -o /ropd

CMD ["/ropd"]
