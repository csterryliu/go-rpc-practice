package main

import (
        "fmt"
        "net/rpc"
)

type Args struct {
        A int
        B int
}

func main() {
        args := Args{5, 2}
        var reply int
        
        client, err := rpc.Dial("tcp", "localhost:3000")  // connet to a rpc server
        checkError(err)
        
        err = client.Call("Arith.Multiply", args, &reply)
        checkError(err)
        fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)
        
        err = client.Call("Arith.Add", args, &reply)
        checkError(err)
        fmt.Printf("%d + %d = %d\n", args.A, args.B, reply)
        
}


func checkError(err error) {
        if err != nil {
             fmt.Println(err)
             return
        }
}