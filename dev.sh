#!/usr/bin/env bash

GO111MODULE=on GOPROXY=https://goproxy.io,direct go build

./rop-back-neo -db_user rop -db_passwd rop_pass -db_addr localhost:5432 -db_name rop
