package main

import (
	//这里导入的是路径 操作系统的路径
	"firstStart/lib"
	"flag"
)
var name string

func init(){
	flag.StringVar(&name, "name", "everyone", "The greeting object.")



}

func main(){



	flag.Parse()
	//但是使用时是 该实体声明的包名下的实体
	lib5.Hello(name)



	//fmt.Printf("Hello, %s!\n", name)
}

