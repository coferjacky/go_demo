package main

import "fmt"

//设置全局变量
var (
	x = 100
	y = "good"
	z = 80.90
)
var block="alibaba"
func main() {
	var x = 200
	fmt.Println("x=", x, "y=", y, "z=", z)

	//以下演示变量重名
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)


}
