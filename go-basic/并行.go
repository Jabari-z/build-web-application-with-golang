package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Print("这里是main（） 开始的地方\n")
	go longWait()
	go shortWait()
	fmt.Println("挂起main（）")
	time.Sleep(10*1e9)
	fmt.Print("这里是main（） 结束的地方\n")
}
func longWait(){
	fmt.Println("开始 long ")
	time.Sleep(5*1e9)
	fmt.Println("结束 long ")

}

func shortWait(){
	fmt.Println("开始 short ")
	time.Sleep(2*1e9)
	fmt.Println("结束 short ")
}

