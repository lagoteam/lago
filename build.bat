@echo off
@REM 开始打包
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build --ldflags "-extldflags -static" -o main .
