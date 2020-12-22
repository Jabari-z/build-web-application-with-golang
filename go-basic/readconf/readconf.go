package main

import "fmt"
import "gopkg.in/gcfg.v1"

func main() {
	config := struct {
		Section struct{
			Enabled bool
			Path string
		}
	}{}
	err:=gcfg.ReadFileInto(&config,"conf.ini")
	if err!=nil{
		fmt.Println("fail")
	}
	fmt.Println("[][][][]")
	fmt.Println(config.Section.Path)
}
