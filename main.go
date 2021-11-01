package main

import (
	"fmt"
	"go_dataStructures/ArrayList"
)

func main()  {
	array := ArrayList.NewArray()
	array.Append("a")
	array.Append("b")
	array.Append("c")
	array.Append("d")
	array.Append("e")
	fmt.Println(array)
	fmt.Println(array.TheSize)

}
func main1()  {
	array := ArrayList.NewArray()
	array.Append(1)
	array.Append(2)
	array.Append(3)
	array.Append(4)
	array.Append(5)
    fmt.Println(array)

}
