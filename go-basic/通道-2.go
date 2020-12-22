package main

import "fmt"

func Count1(ch chan int){
	//通过ch<-1 向对应的channel写入一个数据
	tmp:= <-ch
	tmp++
	ch <- tmp
	fmt.Print(tmp)
}
func main() {
	//定义一个10个channel的数组 名字为chs
	chs:=make([]chan int,10)
	for i:=0;i<10;i++ {
		//把每一个channel分配给gorountine
		chs[i] = make(chan int,i)
		go Count1(chs[i])
	}
	for _,ch:=range chs{
		val := <-ch
		fmt.Print(val)
	}
}