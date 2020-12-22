package main

import "fmt"

type Camera struct {

}

func (c *Camera) TakeAPicture() string{
	return "拍照"
}

type Phone struct {

}
func (p *Phone) Call() string{
	return "🔔响铃"
}

type CameraPhone struct {
	Camera
	Phone
}

func main(){
	cp := new(CameraPhone)
	fmt.Print(cp.TakeAPicture())
	fmt.Print(cp.Call())
}