package main

import "fmt"

func main() {
	//Sprint会把输出保存到sting中
	s1 := fmt.Sprint("zhang")
	name := "bin90"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("bin90")
	fmt.Println("--------")
	fmt.Println(s1, s2, s3)
	//Pprint
	a:=10
	fmt.Printf("%p \n",&a)//指针
	//Scan
	var b int
	fmt.Scan(&b)
	fmt.Println(b)
}