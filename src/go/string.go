package main

import (
	"fmt"
	"go/module"
)

func main() {
	var str = "hello, world"
	// str[0] = 'x'
	// fmt.Print(str)
	for _, value := range str {
		fmt.Println(value)
	}
	//fmt.Print("xx" + strconv.Itoa(3) + str)
	fmt.Print(module.GetName())
}
