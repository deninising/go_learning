package main

import (
	"fmt"
)

func main()  {
	
	pm := new(map[string]int)

	fmt.Printf("type of pm is %T \n",pm)
	fmt.Println("----=----")
	fmt.Println("value of pm is ",pm)
	fmt.Println("address value of pm is ",&pm)

	// panic: assignment to entry in nil map -> new函数也可为引用类型分配内存，但这是不完整创建。以字典（map）为例，它仅分配了字典类型本身（实际就是个指针包装）所需内存，并没有分配键值存储内存，也没有初始化散列桶等内部属性，因此它无法正常工作。
	m := *pm
	m["name"] = 1

}