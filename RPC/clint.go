package main

import (
	"fmt"
	"RPC/rpcservice"
	"log"
	"net/rpc"
)

func main() {
	// 拨号 DialHTTP connects to an HTTP RPC server at the specified network address listening on the default HTTP RPC path.
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// 请求消息
	stringReq := &rpcservice.StringRequest{"A", "B"}
	// Synchronous call
	var reply string
	err = client.Call("StringService.Concat", stringReq, &reply)
	if err != nil {
		log.Fatal("StringService error:", err)
	}
	fmt.Printf("StringService Concat : %s concat %s = %s\n", stringReq.A, stringReq.B, reply)
	// 请求消息 2
	stringReq = &rpcservice.StringRequest{"ACD", "BDF"}
	call := client.Go("StringService.Diff", stringReq, &reply, nil)
	_ = <-call.Done
	fmt.Printf("StringService Diff : %s diff %s = %s\n", stringReq.A, stringReq.B, reply)

}