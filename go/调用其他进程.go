package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//简单调用
	dateCmd := exec.Command("date")
	dateOut,err:=dateCmd.Output()
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(dateOut))

	//
}
