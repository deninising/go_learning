package main


func main()  {
	// 常规用法
	println("常规用法")
	for i:=0 ;i<=5;i++ {
		println("i =",i)
	}

	x:=5
	println("相当于 while(x > 0)")
	// 相当于 while(x > 0)
	for x>0 {
		println("x =",x)
		x--
	}

	println("相当于while(true)")
	// 相当于while(true)
	for {
		if x<5{
			x++
			println("x =",x)
		}else{
			break	
		}
	}


}