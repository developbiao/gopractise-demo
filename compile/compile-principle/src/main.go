package main

import (
    "fmt"
    "gopractise-demo/compile/compile-principle/src/util"
)

func main() {
    fmt.Println("Hello Address!")

    localIp, err := util.GetLocalIPv4Address()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Local IP: %s\n", localIp)
}

