package main

import "fmt"
type N int
// 值传递
func(this N) toString() string{
	return fmt.Sprintf("%#x",this)
}

// 引用传递
func(this *N) pointerRecieve() {
	*this++ 
  fmt.Printf("this: %p, %v\n",this,*this)
	fmt.Println("the address of the pointer reciever input this is ",this)
}

// 值传递/被复制
func(this N) valueRecieve() {
	this++ 
  fmt.Printf("this: %p, %v\n", &this,this)
	fmt.Println("the address of the value reciever input this is ",&this)
}

func main()  {
 	var inputNumber N =100
 	fmt.Println(inputNumber.toString())

	 // 原始对象地址
	 fmt.Println("the source object address is ", &inputNumber)
	 // 值传递地址
	 inputNumber.valueRecieve()
	 // 地址传递
	 inputNumber.pointerRecieve()
}