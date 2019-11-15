package main

import (
	"fmt"
	"os"
)

// os 包获取操作系统的一些信息，os.Args[0] 拿到命令行输入参数。Args变量为切片类型
func main() {
	fmt.Println(os.Args[3])
	//fmt.Print(os.Args[0], os.Args[1])
	//var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//sep = " "
	//}
	//fmt.Println(s)

}
