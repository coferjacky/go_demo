package main

import (
	"fmt"
	"unsafe"
)

//基本数据类型int

func main(){
	var i int8=-128
	var j int=-9223372036854775808
	fmt.Println(i)
	fmt.Println(j)
	x:=10000
	//下列语句可以输出变量的类型以及占用的字节数％
	fmt.Printf("x的类型%T,占用字节数%d",x,unsafe.Sizeof(x))

}
