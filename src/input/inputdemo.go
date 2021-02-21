package main

import "fmt"
//scanln 获取一行值
//scanf 是获取一定格式的输入
func main(){
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)

	fmt.Println("请输入是否通过考试")
	fmt.Scanln(&isPass)

	fmt.Printf("名字是%v \n 年龄是%v\n 薪水是%v \n 是否通过考试%v \n",name,age,sal,isPass)


	//方法二 按照指定格式输入
	fmt.Scanf("%s %d %f %t",&name,&age,&sal,&isPass)
	fmt.Printf("名字是%v \n 年龄是%v\n 薪水是%v \n 是否通过考试%v \n",name,age,sal,isPass)

}
