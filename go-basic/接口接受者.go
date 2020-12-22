package main

import "fmt"

type user struct {
	name  string
	email string
}

type notifier interface {
	notify()
}

//user 的指针类型 user类型没有实现notify方法
func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// 任何实现 notifier接口的值都可以传入方法
func sendNotification(n notifier) {
	n.notify()
}
func main() {
	u := user{"Bill", "bill@email.com"}

	sendNotification(&u)
}
