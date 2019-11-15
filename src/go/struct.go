package main

import (
	"fmt"

	"github.com/pquerna/ffjson/ffjson"
)

type NewsModel struct {
	Id    int
	Title string
}

func main() {
	news := NewsModel{110, "hello"}

	res, err := ffjson.Marshal(news)

	if err != nil {
		fmt.Print(err)
	}

	// 得到是字节数组，所以还有转为string

	fmt.Println(string(res))

}
