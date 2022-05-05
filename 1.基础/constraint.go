package main

import (
	"fmt"
	"unsafe"
)


func main() {
	//cons功能
	con()

	//iota
	iot()
}

//https://blog.csdn.net/buptwcx/article/details/107784638
/*
unsafe.Sizeof返回的是数据类型的大小，而string在Go中并不是直存类型，它是一个结构体类型
type StringHeader struct {
        Data uintptr  //无符号整型，用于存放指针
        Len  int
}
在64位系统上uintptr和int都是8字节，加起来就16了
*/
func con() {
	const (
		a = "abcdeaaaa"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	fmt.Println(a,b,c)
}

// iota是go语言的常量计数器，只能在常量的表达式中使用。
//https://zhuanlan.zhihu.com/p/368595745
func iot() {
	const (
		i = 1 << iota
		j = 3 << iota
		k 
		l
	)
	fmt.Println(i,j,k,l)
}
