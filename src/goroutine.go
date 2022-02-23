package main

import (
	"fmt"
	"runtime"
)

func main()  {
	go say("world") // 通过关键字go就启动了一个goroutine。
	say("hello")    // 当前routine
}


func say(s string)  {
	for i:=0;i<5;i++{
		runtime.Gosched() // runtime.Gosched()表示让CPU把时间片让给别人，下次某个时候继续恢复执行该goroutine。
		fmt.Println(s)
	}
}