package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	fmt.Println(rpc.DialHTTP("tcp", "127.0.0.1:1231"))

}
