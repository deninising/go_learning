package main

import "fmt"

type flags byte

const(
	read flags = 1 << iota
	write 
	exec
)


func main()  {
	fmt.Printf("value of read is %b \n",read)
	fmt.Printf("value of write is %b \n",write)
	fmt.Printf("value of exec is %b \n",exec)

	f := read | exec | write

	fmt.Printf("value of f is %b \n",f)
}