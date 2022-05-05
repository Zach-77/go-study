package main

import "fmt"

func main() {

// [1] if .. else ..
	var test_if int = 95
	fmt.Printf("\n---case1: simple type---\n")
	if test_if >= 90 {
		fmt.Printf("优秀 \n ")
	} else if test_if < 90 && test_if >= 80 {
		fmt.Printf("良好 \n ")
	} else {
		fmt.Printf("一般 \n ")
	}
	
// [2] switch
	/*
		1.普通型
	*/
	fmt.Printf("\n---case2: simple type---\n")
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90 : grade = "A"
	case 80 : grade = "B"
	case 50,60,70 : grade = "C"
	default: grade = "D"
	}

	switch {
	case grade == "A" :
		fmt.Printf("%d :优秀\n", marks )
	case grade == "B" :
		fmt.Printf("%d :良好", marks )
	case grade == "C" :
		fmt.Printf("%d :及格", marks )
	case grade == "F" :
		fmt.Printf("%d :不及格", marks)
	default:
		fmt.Printf("退学吧")			
	}
	fmt.Printf("你的等级是%s \n", grade)


	/*
		2. Type Switch型
		switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	*/
	fmt.Printf("\n---case3: type switch---\n")
	var x interface{}
	
	switch i := x.(type) {
	case nil:
		fmt.Printf("X的类型是: %T", i )
	case int:
		fmt.Printf("X的类型是: int" )
	case float64:
		fmt.Printf("X的类型是: float64" )
	case func(int) float64:
		fmt.Printf("X的类型是: func(int) float64" )
	case bool, string:
		fmt.Printf("X的类型是:  bool or string" )
	default:
		fmt.Printf("未知型")
	}
	fmt.Printf("\n")

	/*
		3. fallthrough
		使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	*/
	fmt.Printf("\n---case4: fallthrough---\n")
	switch {
    case false:
            fmt.Println("1、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("2、case 条件语句为 true")
            fallthrough
    case false:
            fmt.Println("3、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("4、case 条件语句为 true")
    case false:
            fmt.Println("5、case 条件语句为 false")
            fallthrough
    default:
            fmt.Println("6、默认 case")
    }

	/*
		4. select 类型
		select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。
		select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。

		select {
			case communication clause  :
				statement(s);      
			case communication clause  :
				statement(s);
				// 你可以定义任意数量的 case
			default : // 可选 
		statement(s);
		}
		*/
	fmt.Printf("\n---case5: select---\n")
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
		case i1 = <-c1:
			fmt.Printf("received ", i1, " from c1\n")
		case c2 <- i2:
			fmt.Printf("sent ", i2, " to c2\n")
		case i3, ok := (<-c3):  // same as: i3, ok := <-c3
			if ok {
				fmt.Printf("received ", i3, " from c3\n")
			} else {
				fmt.Printf("c3 is closed\n")
			}
		default:
			fmt.Printf("no communication\n")
	}   
}