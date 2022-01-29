package main

func main()  {
	x:= -100
	switch  {
		case x<0:
			println("x < 0")
		case x>0:
			println("x > 0")
		default:
			println("x = 0")
	}
}