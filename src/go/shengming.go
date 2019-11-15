package main

//go语言 有四个声明 关键字
// var 变量声明
//const 常量声明
//type 类型声明  声明并实例化 一种特殊类型 如自定义的类型(Tyoe A int64)  或者（struct和interface）type A struct
//func 函数声明
//函数内部的变量名字必须要先声明 才能使用

func main() {
	// var 变量名字 变量类型 = 表达式
	//变量类型 或者 =表达式可以省略其中一个
	//如果省略变量类型  系统会根据初始化的表达式来推导类型
	//如果省略的是 =表达式 那么会根据类型来给声明的变量初值  int为0 string为“ ” bool为false （切片 指针 map chan）为nil

	//var str string -> str的值为" "
	//var b =true  -> b为bool类型

	//简短 变量的声明和初始化  :=
	// 在main（）函数执行之前 包级别的变量（全局变量已经初始化完成），局部变量要等到函数执行到定义变量那块 再初始化。
	//os.Open(name)会返回2个变量 我们可以用 f，err：= 方式声明并初始化函数的返回的 f,err 这两个变量自动识别返回类型
	// f, err := os.Open(name)
	// if err != nil {
	// 	return err
	// }

	// in, err := os.Open(infile)
	// out, err := os.Create(outfile)
	// 第二个err只有赋值操作

	//  in, err := os.Open(infile)
	//  in, err := os.Create(outfile)
	//编译会出错 in, err = os.Create(outfile) 把下面的改成这样
}
