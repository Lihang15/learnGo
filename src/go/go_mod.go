

一，go 语言本地包如何相互引用
以a.go 引入b.go 中Helloworld方法为例子,函数名字大写其他文件才可用
project/main/a.go
            /b.go
		/mypkg/c.go
		go.mod
		go.sum	

文件夹名字就是包的名字 a.go的内容头 就是package main

a.go 和b.go 在同一个文件夹下，a里面吗可以直接执行Helloworld()函数，不用import

a.go 和c.go  

a.go
package main

import "project/mypkg"

func main() {
    mypkg.Hello() // 调用另一个包里的函数
}




二，配置go mod  用于管理和使用别人写好的三方依赖包


1.  在项目根目录下  go mod init 模块名字  根目录下生成 go.mod 文件    如果你的项目作为三方包给别人用，通过这个模块名字引入

2.  有依赖包包的情况下  执行main.go 会自动拉取包  并生成 go.sum文件 包会被拉取到go默认路径下缓存

3.  go mod vendor  会在根目录生成vendor目录 并把GOPATH/pkg/mod下的包的包复制过来



三，go.mod 升级包的版本，步骤：

直接修改 go.mod 中包的版本，GoLand 会自动下载和更新包
可以执行命令：go mod tidy，会根据代码里引用的包，自动进行包的整理
如果需要同步到 vendor 文件夹，执行命令：go mod vendor
运行：sudo go run main.go，看能否编译通过；如果编译有报错，进行相应包的调整。

注：直接在项目目录下执行go get 也可以拉去所有依赖包  前提下必须目录有go.mod文件