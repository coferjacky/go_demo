package main

import  "fmt"
//import . "fmt" 那么就会让这个“fmt”包中公开的程序实体，被当前源码文件中的代码，视为当前代码包中的程序实体。

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
		//这里block可以是任意的类型，不影响编译
	block := 100
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %d.\n", block)


}
