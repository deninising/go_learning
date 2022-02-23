package main

import "fmt"

type N int

func(n N)test(){
	fmt.Printf("test.n: %p, %d\n", &n, n)
}

func main() {

	var n N=25
  fmt.Printf("main.n: %p, %d\n", &n, n) 
	fmt.Println("---------=-----------")
	// method express 复制给f1,并以值传递的方式使用
	f1 := N.test  // N.test(n)
	f1(n)
	// method express 复制给f1,并以指针传递的方式使用
	f2 := (*N).test // (*N).test(&n)    注意: *N 须使用括号，以免语法解析错误
	f2(&n)     // 按方法集中的签名传递正确类型的参数
}