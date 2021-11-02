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
	array.Append("f")
	array.Append("g")

	for iterator := array.Iterator(); iterator.HaxNext(); {
		item,_ := iterator.Next()
		if item == "e" {
			iterator.Remove()
		}
		fmt.Println(item)
	}
    fmt.Println("*******")
	fmt.Println(array)
}
func main3()  {
	var array ArrayList.List = ArrayList.NewArray()
	array.Append("a")
	array.Append("b")
	array.Append("c")
	for i := 0; i < 10; i++ {
		array.Insert(2,"x5")
		fmt.Println(array)
	}
}
func main2()  {
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
