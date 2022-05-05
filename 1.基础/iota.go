package main

import (
  "fmt"
)

type OrderStatus int

//reffer: https://zhuanlan.zhihu.com/p/368595745

const (
  Cancelled OrderStatus = iota + 1 //订单已取消 1
  NoPay //未支付 2
  PendIng // 未发货 3
  Delivered // 已发货 4
  Received // 已收货 5
)

//公共行为 赋予类型 String() 函数，方便打印值含义
func (order OrderStatus) String() string {
  return [...]string{"cancelled", "noPay", "pendIng", "delivered", "received"}[order-1]
}

//创建公共行为 赋予类型 int 函数 EnumIndex()
func (order OrderStatus) EnumIndex() int {
  return int(order)
}

func main() {
  var order OrderStatus = Received
  fmt.Println(order.String()) // 打印:received
  fmt.Println(order.EnumIndex()) // 打印:5
}