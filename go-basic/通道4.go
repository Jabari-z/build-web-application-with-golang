package main

import (
	"fmt"
	"time"
)

func PrintCount(c chan int){
	num:=0
	fmt.Println("for 前")
	for true{
		fmt.Println("        :num = <-c 前"	,num)
		num = <-c
		fmt.Println("in for：",num)
		fmt.Println("        :print后\n----------")
	}
	fmt.Println("for 后")
}
func main() {
	c:=make(chan int) 			//channel
	a:=[]int{8,6,-1,7,5,3,0,-1,9}  //数组
	
	go PrintCount(c)   //启动goroutine

	//把数组中的值送给channel
	for _,v :=range a{
		c <- v
	}
	time.Sleep(time.Microsecond*100) // 主程序等待事件： 时间长一点， PrintCount就会在主程序之前结束
	fmt.Println("----------End")
}