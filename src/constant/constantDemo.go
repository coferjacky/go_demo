package main

import "fmt"

const (
	pi=31.3
	jack=32
	book

)
const (

	_=iota
	Kb=1<<(10*iota)
	Mb=1<<(10*iota)
	Gb=1<<(10*iota)
	TB=1<<(10*iota)
	PB=1<<(10*iota)
)
func main()  {
	fmt.Println(pi)
	fmt.Println(jack)
	fmt.Println(book)

	fmt.Println(Kb)
	fmt.Println(Mb)
	fmt.Println(Gb)
	fmt.Println(TB)
	fmt.Println(PB)
}