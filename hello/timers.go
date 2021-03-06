package main

import "fmt"
import "time"

func main() {
    // Waiting 3 second
    timer1 := time.NewTimer(time.Second * 3)
    <-timer1.C
    fmt.Println("Timer 1 expired")

    timer2 := time.NewTimer(time.Second)
    go func() {
       <-timer2.C     
       fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 has been stopped")
    }
    

}
