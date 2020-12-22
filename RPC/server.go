package main

import (
	"RPC/rpcservice"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	stringService := new(rpcservice.StringService)
	//注册
	registerError := rpc.Register(stringService)
	//注册错误处理
	if registerError != nil {
		log.Fatal("Register error: ", registerError)
	}
	// HandleHTTP registers an HTTP handler for RPC messages
	// to DefaultServer on DefaultRPCPath and a debugging handler
	// on DefaultDebugPath.
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "127.0.0.1:1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
