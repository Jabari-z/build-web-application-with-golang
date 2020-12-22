package test

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T){
	name:=getName()
	if name!="world"{
		t.Error("不是期望值")
	} else{
		fmt.Println("test")	
	}
}