package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s : %v ", os.Args[1], err)
		os.Exit(1)
	}
	str, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Print("err")
		os.Exit(1)
	}
	fmt.Printf("%s", str)
}
