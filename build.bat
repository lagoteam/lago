@echo off
chcp 65001

echo "开始打包..."

set CGO_ENABLED=0
set GOARCH=amd64

set "input=0"
set /p "input=请输入操作系统数字编号 [0-Linux(默认), 1-Windows] :"

if "%input%"=="1" (
    echo "Windows"
	set GOOS=windows
    go build --ldflags "-extldflags -static" -o main.exe .
) else (
    echo "Linux"
	set GOOS=linux
    go build --ldflags "-extldflags -static" -o main .
)

echo "打包成功"

@REM pause