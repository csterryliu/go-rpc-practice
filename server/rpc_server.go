package main

import (
        "net"
        "net/rpc"
        "fmt"
)

type Args struct {
        A int
        B int
}

type Arith int   // you can consider it as a service of RPC
func (a *Arith) Add(args *Args, reply *int) error {  // a method of service
        *reply = args.A + args.B
        return nil
}
func (a *Arith) Multiply(args *Args, reply *int) error {  // a method of service
        *reply = args.A * args.B
        return nil
}


func main() {
        service := new(Arith)  // this service contains 2 methods
        rpc.Register(service)  // register this service to DefaultServer. the name of this service will be its concrete type
        
        tcpAddr, err := net.ResolveTCPAddr("tcp", ":3000")
        checkError(err)
        listener, err := net.ListenTCP("tcp", tcpAddr)
        checkError(err)
        
        // Let DefaultServer starts to wait for requests
        fmt.Println("start running...")
        rpc.Accept(listener)
        
}


func checkError(err error) {
        if err != nil {
             fmt.Println(err)
             return
        }
}