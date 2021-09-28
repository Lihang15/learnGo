package main

import (
	"fmt"
)

//go run 问题 如下
//用go run helloworld.go 可以可以编译一个或者多个以.go为结尾得go文件，并且链接库文件，生成2进制可执行文件。
//生成的可执行文件由于是静态编译，所以不用担心依赖的库升级问题，拿到go环境，随便运行。

//构建go工程问题 如下
//一般写go工程代码要把工程目录添加到GOPATH  然后go get 就可以向GOPATH下的工程拉取别人写的3方go源代码 会拉取到GOPATH/src目录 你就可以用了
//go 1.11版本以后 包含1.11 一般用 go.mod 构建项目 不需要把项目放到GOPATH下 

//go语言的包组织问题 如下
//go语言通过 package组织包 方便别的地方调用  一般你的包的名字 和你所在文件夹名字相同 别的包调用你的包函数时候你就直接导入
//你这个文件路径就可以了 然后用包名字去 . 你的函数。
func main() {
	fmt.Print("hello world")
}
