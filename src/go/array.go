package main

func main() {
	//var a [3]int 声明并初始化数组为0 0 0 长度为三

	//遍历数组
	// for index, value := range a {
	// 	fmt.Println(index, value)
	// }

	//fmt.Print(a[0], a[1])

	// 初始化数组自定义值
	// var a [3]int = [3]int{1, 2, 3}
	// a[0] = 100 //字符串和数组都为不可变类型  数组可以通过下标重赋值  字符串不行
	// for index, value := range a {
	// 	fmt.Println(index, value)

	// }

	//没有固定长度的特殊数组  也就是切片  一般都用切片。
	// a := []int{1, 2, 3, 4, 5, 6}   定义切片
	// a := [6]int{1, 2, 3, 4, 5, 6}  定义数组
	// fmt.Print(a[0:])
	// arr := make([]int, 0)
	// arr = append(arr, 10)
	// fmt.Print(arr[0:])

}
