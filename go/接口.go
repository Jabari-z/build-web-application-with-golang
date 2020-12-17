package main

import "fmt"

/*
接口定义
type Namer interface{
	Method1(param_list) return_type
	Method1(param_list) return_type
	Method1(param_list) return_type
}
 */

type error interface {
	Error() string
}

type RPCError struct {
	Code    int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func main() {
	var rpcErr error = NewRPCError(400, "unknown err") // typecheck1
	err := AsErr(rpcErr) // typecheck2
	println(err)
}

func NewRPCError(code int64, msg string) error {
	return &RPCError{ // typecheck3
		Code:    code,
		Message: msg,
	}
}

func AsErr(err error) error {
	return err
}