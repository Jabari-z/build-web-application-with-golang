package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}){
	t:=reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n",t.Name(),t.Kind())   // 注意这里 type是类型名 kind是底层的名字
}
func main(){
	var x float64= 3.1
	v := reflect.TypeOf(x)
	fmt.Println(v.Kind())
	fmt.Println(reflect.TypeOf(x))

	type person struct{
		name string
		age int
	}

	type book struct{
		title string
	}

	var d = person{name:"小王子",age:18}
	var e =book{title:"Go language"}
	reflectType(d)
	reflectType(e)
	fmt.Println("不窜爱的结构体成员",reflect.ValueOf(e).FieldByName("abv").IsValid())
	// IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
	name := "bingo"
	valOfName:=reflect.ValueOf(&name)
	//valOfName.Elem().Set(reflect.ValueOf("xiaohong"))
	elem := valOfName.Elem()
	elem.Set(reflect.ValueOf("xiaohong"))

	fmt.Println(name)
}
