package main

import (
	"fmt"
	"net"
)

type client chan string

var (
	message  = make(chan string) //装客户端发来的数据
	entering = make(chan client)
	leaving  = make(chan client)
)

func main() {
	//注册tcp服务 开放9090端口 拿到监听对象
	Listener, err := net.Listen("tcp", ":9090")

	if err != nil {
		fmt.Print("获取监听者对象失败")
	}
	//开始监听
	for {
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Print("没能拿到客户连接继续去拿")
			continue
		}
		// //如果拿到了客户端ip地址 存到数据库
		// fmt.Print("拿到客户端ip的", conn.RemoteAddr().String())
		// //先往数据库里存上。
		go handleConn(conn)

	}

}

func handleConn(conn net.Conn) {

}

// //拿到客户连接后的处理
// func hander(conn net.Conn) {
// 	//定义字节数组为了存储客户端来的数据
// 	acceptClientData := make([]byte, 1024)
// 	msg, err := conn.Read(acceptClientData)
// 	if err != nil {
// 		fmt.Print("读取客户端数据失败")
// 		return

// 	}
// 	fmt.Print("输出来自客户端的数据", string(acceptClientData[:msg]))
// }
