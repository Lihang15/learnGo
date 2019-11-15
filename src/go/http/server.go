package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hander)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func hander(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "urlpath is %s", req.URL.Path)

}
