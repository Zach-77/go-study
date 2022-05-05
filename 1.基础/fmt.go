package main

import (
	"fmt"
)

func main() {
	//%d 表示整数型数字，%s表示字符串
	var stockcode = 123
	var enddate = "2022-04-24"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url,stockcode,enddate)
	fmt.Println(target_url)

}