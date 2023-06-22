#!/bin/sh

echo "脚本开始运行..."

echo "下载并解压Golang..."
wget -c https://dl.google.com/go/go1.19.6.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local

echo "配置Golang环境变量..."
export PATH=$PATH:/usr/local/go/bin

sudo tee -a ~/.profile << EOF
export GOROOT=/usr/local/go
export GOPATH=$HOME/gowork
export PATH=\$PATH:\$GOROOT/bin:\$GOPATH/bin
EOF

echo "重新加载环境变量..."
sudo source ~/.profile

echo "查看Golang版本号..."
go version

echo "设置加速镜像..."
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

echo "查看Golang环境变量..."
go env

echo "下载项目依赖..."
#go mod download
go mod tidy
#go clean -cache
#go clean -modcache

#go get -v -u github.com/cosmtrek/air
#go get -u github.com/kataras/rizla

echo "正在编译应用..."
go env -w CGO_ENABLED=0
#go env -w GOOS=linux
#go env -w GOARCH=amd64
go build --ldflags "-extldflags -static" -o main .

echo "正在给文件添加权限..."
sudo chmod a+x ./*
sudo chown -R www . && sudo chmod 755 -R ./

echo "脚本运行完成"
