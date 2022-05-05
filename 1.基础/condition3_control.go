package main

import (
	"fmt"
)

func main() {
/*	case 1:
	theme: break
*/
	fmt.Printf("\n---case1 ---\n")
	
	//don't use lable
	fmt.Printf("---- [break] don't use break lable -----\n")
	for i := 1; i <= 10 ; i++ {
		fmt.Printf("i: %d\n",i)
		for i2 := 100; i2 <= 200; i2++ {
			fmt.Printf("i2: %d \n", i2)
			break
		}
	}
	//fmt.Printf("i: %d\n",i)

	// use lable
	fmt.Printf("---- [break] use break lable -----\n")
	test:
		for i := 1 ; i<= 3 ; i++ {
			fmt.Printf("i: %d\n",i)
			for i2 := 100; i2 <= 200 ; i2++ {
				fmt.Printf("i2:%d\n", i2)
				break test
			}
		}

/*	case 2:
	theme: continue 
*/
	fmt.Printf("\n---case2 ---\n")

	//don't use lable
	fmt.Printf("---- [continue] don't use break lable -----\n")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n",i)
		for i2 := 11; i2 <= 13; i2++{
			fmt.Printf("i2 :%d\n",i2)
			continue
			}
		}

	// use lable
	fmt.Printf("---- [continue] use break lable -----\n")
	re:
		for i := 1; i <= 3; i++ {
			fmt.Printf("i: %d\n",i)
			for i2 := 11; i2 <= 13; i2++{
				fmt.Printf("i2 :%d\n",i2)
				continue re
			}


/*	case 3:
	theme: goto 
	
	Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
	goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
	但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难
*/
	fmt.Printf("\n---case2 ---\n")
	
	LB : for a < 30 {
		if i == 15 {
			a = a + 1
			goto LB
		}
		fmt.Printf("a的值为: %d\n",a)
		a++
	}

	}
}