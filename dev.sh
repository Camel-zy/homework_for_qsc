#!/usr/bin/env bash

GO111MODULE=on GOPROXY=https://goproxy.io,direct go build

if [[ $? -ne 0 ]]; then
    exit 1
fi

./rop-back-neo -db_user rop -db_passwd rop_pass -db_host localhost -db_port 5432 -db_name rop # -db_init
