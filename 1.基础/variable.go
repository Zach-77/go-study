package main

import "fmt"

var x, y int 

var (  //这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123,"hello"

//这种不带声明格式的只能在函数体中出现
// g,h := 123,"hello"

func main(){
	g,h := 123,"hello"
	println("test1:",x,y,a,b,c,d,e,f,g,h)

	_,numb,strs := numbers()  //只获取函数返回值的后两个, _ 为抛弃值
	fmt.Println("test2:",numb,strs)
}

//一个可以返回多个值的函数
func numbers()(int,int,string){
	a1 , b1 , c1 :=1 , 2 , "str"
	return a1,b1,c1
}