package main
import "fmt"

func main() {
    const lightSpeed = 29972 // km/s
    var distance = 56000000 // km
    fmt.Println(distance / lightSpeed, "seconds")
    distance = 401000000
    fmt.Println(distance / lightSpeed, "seconds")
}
