package main

import (
	"fmt"
)

//go run 问题 如下
//用go run helloworld.go 可以可以编译一个或者多个以.go为结尾得go文件，并且链接库文件，生成2进制可执行文件。根据系统默认区分
//生成的可执行文件由于是静态编译，所以不用担心依赖的库升级问题，拿到go环境，随便运行。

func main() {
	fmt.Print("hello world")
}
