package main

import "fmt"

func main(){
	var a byte='a'
	var b byte= 'b'
	var c byte='0'
	var d int='北'
	//转义字符常量
	var ch byte='\a'
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("c=%c\n",a)
	fmt.Printf("d=%c\n",d)
	fmt.Printf("ch=%c",ch)
}

