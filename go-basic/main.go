package main

import "fmt"

//interface
//type Shaper interface {
//	Area() float32
//}
//正方形
type Square struct {
	side float32
}
// 面向对象
func (sq *Square) Area() float32{
	return sq.side*sq.side
}
func main() {
	sq1:=new(Square)
	sq1.side = 5.0
	fmt.Println(sq1.Area())
}

