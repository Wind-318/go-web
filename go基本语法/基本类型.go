package main

import "fmt"

func hello() {
	fmt.Println("hello world")
}

func main() {
	//arr是切片类型，可以装载空接口类型，由于go中所有类型默认都实现了空接口，所以所有类型都可以被装下
	arr := make([]interface{}, 0)
	//放入一个1
	arr = append(arr, 1)
	//放入一个bool类型
	arr = append(arr, true)
	//放入一个函数
	arr = append(arr, hello)
	//放入一个map
	arr = append(arr, make(map[int]string, 0))
	//放入一个channel
	arr = append(arr, make(chan int, 0))
	//打印arr
	fmt.Println(arr)

	/*
		输出：
		[1 true 0xdf57a0 map[] 0xc00004a060]
	*/
}
