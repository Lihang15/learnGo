package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// go 封装的http底层 会为所有的请求分别创建一个携程，为了避免同一时刻，多个线程去改变count出现错误，枷锁。
var lock sync.Mutex
var count int

func main() {
	http.HandleFunc("/h", header)
	http.HandleFunc("/add", hander)
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func hander(w http.ResponseWriter, req *http.Request) {
	lock.Lock()
	count++
	lock.Unlock()
	fmt.Fprintf(w, "request url is %q", req.URL.Path)

}
func counter(w http.ResponseWriter, req *http.Request) {
	lock.Lock()
	fmt.Fprintf(w, "count == %s", count)
	lock.Unlock()
}

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
