package main

import (
	"fmt"
	"reflect"
)

type student struct{
	Name string `json:"name"`
	Score int   `json:"score""`
}

func main() {
	stu1:=student{Name: "xiaowangzi",Score: 90}
	
	t:=reflect.TypeOf(stu1)

	fmt.Println(t.Name(),t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field :=t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	fmt.Println("------")
	if scoreField,ok:=t.FieldByName("Score");ok {
		fmt.Printf("name:%s index:%v type:%v json tag:%v\n",scoreField.Name,scoreField.Index,scoreField.Type,scoreField.Tag.Get("json"))

	}
	printMethod(stu1)
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	fmt.Println("printMethod")
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t,v)
	fmt.Println(t.NumMethod()) // 方法个数
	fmt.Println(v.NumMethod()) // 方法个数
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}

}