package main

import (
	"fmt"
)

/*定义结构体*/
type Circle struct {
	radius float64
}


/*Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给
  定类型的方法属于该类型的方法集。语法格式如下：
*/
func main() {
	var c1 Circle
	c1.radius = 8
	fmt.Printf("半径为%.2f 的圆，面积为:%.2f\n",c1.radius,c1.getArea())
}


//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
 	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

/*
Go 没有面向对象，而我们知道常见的 Java。
C++ 等语言中，实现类的方法做法都是编译器隐式的给函数加一个 this 指针，而在 Go 里，这个 this 指针需要明确的申明出来，其实和其
它 OO 语言并没有很大的区别。

在 C++ 中是这样的

class Circle {
  public:
    float getArea() {
       return 3.14 * radius * radius;
    }
  private:
    float radius;
}

// 其中 getArea 经过编译器处理大致变为
float getArea(Circle *const c) {
  ...
}
在 Go 中则是如下:

func (c Circle) getArea() float64 {
  //c.radius 即为 Circle 类型对象中的属性
  return 3.14 * c.radius * c.radius
*/