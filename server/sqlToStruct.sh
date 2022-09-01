#! /bin/bash

host=192.168.0.125
port=3306
username=root
password=123456
database=monitor
charset=utf8mb4



echo "=======================  开始生成对应的机构体  ======================="

xorm reverse mysql $username:$password@\($host:$port\)/$database?charset=$charset $GOPATH/pkg/mod/github.com/go-xorm/cmd/xorm@v0.0.0-20190426080617-f87981e709a1/templates/goxorm/ ./internal/v1/models/

echo "=======================  开始生成对应的机构体  ======================="