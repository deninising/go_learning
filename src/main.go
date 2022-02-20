package main

import (
	"fmt"
)

var x = 100

func main() {
	// var score int32
	// var string = "hello the world from go"
	// fmt.Println(score, string)
	var a,s=100,"abc"
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(&a,a)

	fmt.Println("=========================================================")

	println(&x,x)       // 全局变量
  
	x:= "abc"          // 重新定义和初始化同名局部变量

	println(&x,x) 
}
