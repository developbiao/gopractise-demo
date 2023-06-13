package main
import "fmt"
import "unsafe"

func main() {
    str :=  bytes2string([]byte("Hello"))
    fmt.Println(str)
}

func bytes2string(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}
