package main

import "fmt"

func main(){
	var a int = 10
	var b int = 20

	fmt.Println("a = ",a," b = ",b)

	// swap
	fmt.Println("main::swap excute...")

	swap(&a,&b)
	fmt.Println("a = ",a," b = ",b)
}


func swap(pa , pb *int) {
	fmt.Println(pa)
	fmt.Println(pb)


	var temp = *pa
	*pa = *pb
	*pb = temp
}