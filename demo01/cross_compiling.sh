#!/bin/bash
# -*- coding:utf-8 -*-
# Windows 下编译 Mac 和 Linux 64位可执行程序
# GOOS：目标平台的操作系统( darwin, freebsd, linux, windows )
# GOARCH：目标平台的体系架构( 386, amd64, arm )
# 交叉编译不支持 CGO 所以要禁用它

go env -w CGO_ENABLED=0
go env -w GOOS=windows
go env -w GOARCH=amd64
go build -o ./target/main main.go
echo 编译成功