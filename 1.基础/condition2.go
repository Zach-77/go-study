package main 

import "fmt"

func main() {

	/*	case 1:
		for init; condition; post { }
	*/
	fmt.Printf("\n---case1 ---\n")
	sum1 := 0
	for i := 0; i <= 10 ; i++ {
		sum1 += i
	}
	fmt.Printf("%d",sum1)




	/* case 2:
		for condition { }
	*/
	fmt.Printf("\n---case2 ---\n")
	
	sum2 := 1
	for ; sum2  <= 10 ; {
		sum2 += sum2
	}
	fmt.Println(sum2)

	//conditon型
	fmt.Printf("\ncase2.1---\n")
	tes := 1
	for tes <= 10{
		tes += tes
	}
	fmt.Println(tes)



	/*	case 3:
		for { }
	*/
	//无线循环

	fmt.Printf("\n---case3 ---\n")

	sum3 := 0
	for {
		sum3++ // 无限循环下去
//		fmt.Println(sum3)		
		if sum3 > 100 {
			break;
		}
	}
	fmt.Println(sum3)

	
	/*
		case 4:
		for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
	*/
	fmt.Printf("\n---case4 ---\n")

	strings := []string{"google", "runoob"}
	for i, s := range strings {
			fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i,x:= range numbers {
			fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}  

	/*
		case 5:
		for 循环嵌套， 求质数
		后面开个多线程算算
	*/
	var i, j int
	for i=2; i < 30; i++ {
	   for j=2; j <= (i/j); j++ {
		  if(i%j==0) {
			 break; // 如果发现因子，则不是素数
		  }
	   }
	   if(j > (i/j)) {
		  fmt.Printf("%d  是素数\n", i);
	   }
	} 
}
