package main

import (
	"os"
	"strconv"
)

//接口
type Stringer interface {
	String() string
}

//结构体
type Celsius float64

//实现接口
func (c Celsius) String() string{
	return strconv.FormatFloat(float64(c),'f',1,64)+"摄氏度"
}

//结构体
type Day int

var dayName = []string{"星期一","星期二","星期三","星期四","星期五","星期六","星期日"}
//实现接口
func (day Day) String() string{
	return dayName[day]
}

func print(args ...interface{}){
	for i,arg := range args{
		if i>0 {os.Stdout.WriteString(" ")}
		switch a:=arg.(type) {
		case Stringer: os.Stdout.WriteString(a.String())
		case int:      os.Stdout.WriteString(strconv.Itoa(a))
		case string:   os.Stdout.WriteString(a)
		default:
			os.Stdout.WriteString("???")

		}
	}
}
func main(){
	print(Day(1),"温度是",Celsius(37.2))
}