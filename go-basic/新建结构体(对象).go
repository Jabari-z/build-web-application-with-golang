package main

import "strings"
type Person struct {
	firstName string
	lastName string
}

func upPersion(p *Person){
	p.firstName = strings.ToUpper(p.firstName)
}

func main(){
	//方式1⃣
	var person Person
	person.firstName = "zhang"
	person.lastName = "jiabin"
	//方式2⃣
	person2 := new(Person)
	person2.firstName = "zhang"
	(*person2).lastName = "jb"
	//方式3⃣
	person3 := &Person{"zhang","jiabin"}
	upPersion(person3)

}