package main

import (
	"errors"
	"fmt"
	"log"
)

func main()  {
	// input := -1
	input := 1
	
	if err := check(input);err == nil {
		input++
		fmt.Println("the value of input++ is ",input)
	}else {
		log.Fatalln(err)
	}
}

// 检测函数
func check(input int)  error{
	if input < 0 {
		return errors.New("input value < 0")
	}
	return nil
}