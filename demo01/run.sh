#!/bin/bash
# -*- coding:utf-8 -*-
# 如果使用code runner执行命令出现乱码, 在终端中输入chcp 65001设置utf-8, 再执行code runner

go build -o ./target/main main.go
./target/main