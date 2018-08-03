package main

import (
	"fmt"
	"net/rpc"
	"log"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}


func main() {
	serverAddress := "localhost:1234"
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args  := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiple", args, &reply)

	if err != nil {
		log.Fatal("arith error:" ,err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	var qot Quotient
	err = client.Call("Arith.Divide", args, &qot)
	if err != nil {
		log.Fatal("Arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, qot.Quo, qot.Rem)

}