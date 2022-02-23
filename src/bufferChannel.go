package main

import "fmt"

/**Go语言也允许指定channel的缓冲大小，
很简单，就是channel可以存储多少元
素。ch:= make(chan bool, 4)，创
建了可以存储4个元素的bool型channel。
在这个channel中，前4个元素可以无阻
塞的写入。当写入第5个元素时，代码将
会阻塞，直到其他goroutine从channel
中读取一些元素，腾出空间。**/
func main()  {
	c := make(chan int,2)

	c <- 1
	c <- 2
	
	fmt.Println(<-c)
	fmt.Println(<-c)

	// fibnacci
	fmt.Println("----------------------=----------------------------------")
	channel := make(chan int,10)
	go fibnacci(cap(channel),channel)
	// 从chan中去读取资源是阻塞状态的，所以 资源生产者如果不停止对channel发送资源，当前循环回一直执行下去。因此若要停止当前循环，需要生产者主动关闭通道close（chan）
	for v := range channel { 
			println("the current fib value is ",v)
	}
}


/**
注： 记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic。
另外，channel不像文件之类需要经常去关闭，只有当你确实没有任何数据发送了，或者你想显式的结束range循环之类的操作。
**/
func fibnacci(number int,c chan int)  {
	x,y := 1,1
	for i := 0;i < number;i++{
		c <- x
		x,y = y, x + y
	}
  close(c)
}