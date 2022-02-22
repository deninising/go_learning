package main

import "fmt"

func main()  {
	
	node1 := node{
		id:1,
	}

	node2 := node{
		id:2,
		next:&node1,
	}

	fmt.Println(node1,node2)


	/**
	var m manager
	m.username = "张三"
	m.age = 18
	m.title = "信息部主管"
	println(m.username,m.age,m.title)
	**/

}

type node struct{
	_ int
	id int
	// 引用字段，或则叫做对象字段
	next *node
}

type user struct{
	username string
	age byte
}

type manager struct{
	user
	title string
}