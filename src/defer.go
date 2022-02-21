package main

import "fmt"

func main() {
	// deferFun()
 	// returnInt()

	deferAndReturn()
}


func deferAndReturn() int {
	defer deferFun()
	return returnInt()
}


func returnInt() int {
	fmt.Println("returnInt excute ...")
	return 0
}


func deferFun(){
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
}