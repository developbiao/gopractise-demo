package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("golang123123bot.com")
	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) {
			if dnsErr.Timeout() {
				fmt.Println("operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("temporary error")
				return
			}
			fmt.Println("Generic DNS error")
			return
		}
	}
	fmt.Println("Resloved address:", addr)
}
