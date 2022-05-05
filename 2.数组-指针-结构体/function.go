package main

import (
	"fmt"
	"math"
)

func main() {
/*
	1.普通值交换
*/
	fmt.Printf("\n---case1 ---\n")

	var a,b string = "男","女"
	fmt.Println(swap(a,b))


/*
	2.引用值传递
*/
	fmt.Printf("\n---case2 ---\n")

	var x,y int = 10,20
	fmt.Printf("值交换前 x:%d ,y:%d \n",x,y)
	swap_point(&x, &y)
	fmt.Printf("值交换后 x:%d ,y:%d \n",x,y)


/*
	3.使用 函数定义后可作为另外一个函数的实参数传入
*/
	fmt.Printf("\n---case3 ---\n")
	/*声明函数变量*/
	getSquareBoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareBoot(64))


/*
	4.闭包
*/
	fmt.Printf("\n---case4-1 ---\n")
	add_func := add(1,2)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())

	fmt.Printf("\n---case4-2 ---\n")
	add_func2 := addV2(1,2)
	fmt.Println(add_func2(4,5))
	fmt.Println(add_func2(1,3))
	fmt.Println(add_func2(2,2))
}

func swap(x , y string) (string,string) {
	return y,x
}

func swap_point(x *int, y *int){
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

// 闭包使用方法
func add(x1, x2 int) func()(int,int)  {
    i := 0
    return func() (int,int){
        i++
        return i,x1+x2
    }
}

// 闭包使用方法，函数声明中的返回值(闭包函数)不用写具体的形参名称
func addV2(x1, x2 int) func(int, int) (int, int, int) {
	i := 0
	return func(x3, x4 int) (int, int, int) {
	  i += 1
	  return i, x1 + x2, x3 + x4
	}
}