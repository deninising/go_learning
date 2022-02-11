package main

import (
	"fmt"
)

func main() {
	dict := make(map[string]int)

	dict["age"] = 10

	fmt.Println(dict["age"])

	msg := make(map[string]string)
	msg["username"] = "张三"
	msg["hobby"] = "游泳"
	fmt.Println(msg["username"])
	fmt.Println(msg["hobby"])
	println(msg["username"], msg["hobby"])

	username, ok := msg["user"]
	println(username, ok)

}