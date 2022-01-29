package main

import "fmt"

func main()  {
	x:=100

  f:= testClosedPg(x)
	f()
}

func testClosedPg(x int) func()  {
	fmt.Println("函数中 x = ",x)
	return 	func()  {
		fmt.Println("闭包中 x = ",x)
	}
}