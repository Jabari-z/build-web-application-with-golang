package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0

func Count(lock *sync.Mutex){
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func main() {
	lock:=&sync.Mutex{}
	for i:=0;i<10;i++{
		// 开启十个协程 共享 counter 变量.10
		go Count(lock)
	}

	for{
		lock.Lock()
		c:=counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}

