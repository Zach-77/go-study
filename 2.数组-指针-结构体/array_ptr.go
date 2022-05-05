package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("交换前a:%d\n",a)
	fmt.Printf("交换前b:%d\n",b)

	sawp(&a,&b)
	fmt.Printf("交换后 a的值 :%d\n",a)
	fmt.Printf("交换后 b的值 :%d\n",b)
}

func sawp(x *int ,y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}