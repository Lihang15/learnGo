
配置go mod

mkdir project 
cd project
project 为根目录

需要设置 GOPATH=xxx目录
设置 GO111MODULE = auto

1.  在项目根目录下  go mod init 模块名字  根目录下生成 go.mod 文件

2.  有依赖包包的情况下  执行main.go 会自动拉取包  并生成 go.sum文件 包会被拉取到 GOPATH/pkg/mod缓存下

3.  go mod vendor  会在根目录生成vendor目录 并把GOPATH/pkg/mod下的包的包复制过来

4. 引入本地包 （go mod init 模块名字）  引入这个名字下的包 模块名字/xxx包

go.mod 升级包的版本，步骤：

直接修改 go.mod 中包的版本，GoLand 会自动下载和更新包
可以执行命令：go mod tidy，会根据代码里引用的包，自动进行包的整理
如果需要同步到 vendor 文件夹，执行命令：go mod vendor
运行：sudo go run main.go，看能否编译通过；如果编译有报错，进行相应包的调整。

注：直接在项目目录下执行go get 也可以拉去所有依赖包  前提下必须目录有go.mod文件