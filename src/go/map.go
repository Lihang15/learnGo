package main

import "fmt"

func main() {
	// mapa := make(map[string]int)
	// mapa["name"] = 1
	// mapa["age"] = 18

	var mapa map[string]int = map[string]int{
		"name": 1,
		"age":  18,
	}
	// for key, value := range mapa {
	// 	fmt.Println(key, value)
	// }
	for value := range mapa {
		fmt.Println(value)
	}
}
