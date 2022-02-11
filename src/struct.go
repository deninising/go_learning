package main

func main()  {
	
	var m manager
	m.username = "张三"
	m.age = 18

	m.title = "信息部主管"
	
	println(m.username,m.age,m.title)

}

type user struct{
	username string
	age byte
}

type manager struct{
	user
	title string
}