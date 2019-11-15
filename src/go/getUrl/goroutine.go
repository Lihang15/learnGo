package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// fetch 抓取url 和内容 放到通道  main 并发的从通道拿到所有携程的数据
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	//拿数据的时候可能会阻塞 如果chan 里没有数据。
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	end := time.Since(start).Seconds()
	fmt.Printf("main routine run %.2fs", end)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbyte, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	end := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", end, nbyte, url)
}
