package main

import "fmt"

func main()  {
	var i="你好中国110   bagadsfdsa"
	//使用反引号来显示源代码
	var code=`
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
	

`
	//字符串拼接方式
	var str="hello"+"world"
	str+="hahah!"
	str4:="hello"+
		"hellow"+
		"workld"


	fmt.Println(i)
	fmt.Println(code)
	fmt.Println(str)
	fmt.Println(str4)

}
