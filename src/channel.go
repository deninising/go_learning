package main

import "fmt"

	/**默认情况下，channel接收和发送数据都是阻塞的，
	除非另一端已经准备好，这样就使得Goroutines同步
	变得更加简单，而不需要显式的lock。所谓阻塞，也
	就是如果读取（value:= <-ch），它将会被阻塞，直
	到有数据接收。另外，任何发送（ch<-5）将会被阻塞，
	直到数据被读出。无缓冲channel是在多个goroutine
	之间同步很棒的工具。**/
func main()  {
	numbers := []int{7,2,8,-7,6,-1,0,12}
	channel := make(chan int)	

	var leftHalfResult int
	var rightHalfResult int
	go sum(numbers[:len(numbers)/2],channel) // channel 遵循一次接口 一次发送的通讯准则，并且多是阻塞通讯
	go sum(numbers[len(numbers)/2:],channel)

	recieve(&leftHalfResult,channel)
	recieve(&rightHalfResult,channel)
	fmt.Println(leftHalfResult,rightHalfResult)
}

func sum(numbers []int,c chan int)  {
	sum := 0
	for _,v := range numbers {
		sum = sum + v
	}

	fmt.Println("channel准备接受数据:",sum)
	c <- sum // 将累加结果发动到channel  c 中去
	fmt.Println("channel准备接受完成:",sum)
}

func recieve(reciever *int,c chan int){
	fmt.Println("channel准备发送数据",reciever)
	*reciever = <- c
	fmt.Println("channel准备发送完成",reciever)
}