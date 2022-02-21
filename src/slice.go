package main

import "fmt"

func main()  {

// 申明
 x:=	make([]int,0,5)
 for i := 0; i < 8; i++ {
	 x=append(x, i)
 }
 fmt.Println(x)

 slice1:= []int{1,2,3}
// 遍历
 for idx,v := range slice1{ 
	fmt.Println("idx = ",idx,"value = ",v)
 }
// 开辟具有三个元素为0 ，容量为5的切片
 number := make([]int,3,6)
 fmt.Println("len = ",len(number),"cap = ",cap(number),"number = ",number)
 
 // 动态追加元素
 number = append(number, 1)
 fmt.Println("len = ",len(number),"cap = ",cap(number),"number = ",number)
 // 自动扩容为原来cap的两倍
 number = append(number, 2)
 number = append(number, 3)
 number = append(number, 4)
 fmt.Println("len = ",len(number),"cap = ",cap(number),"number = ",number)

 // copy（）与截取
}