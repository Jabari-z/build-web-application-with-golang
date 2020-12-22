package test

import "fmt"
func getName () string {
	return "world"
}

func main() {
	name:=getName()
	fmt.Println("hello",name)
}
