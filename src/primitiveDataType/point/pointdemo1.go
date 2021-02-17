package main
import "fmt"

func main(){
	var x int=89
	fmt.Println("x地址 ",&x)
	//ptr本身值为一个地址值
	var ptr *int =&x
	fmt.Println("ptr=%v",ptr)
	//ptr也指向一个空间 我们看看这个空间地址
	fmt.Println("ptr地址是",&ptr)

	fmt.Println("取出ptr指针指向的地址值",*ptr)
}
