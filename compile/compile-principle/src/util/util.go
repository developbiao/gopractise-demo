package util

import (
    "net"
    "errors"
)

func GetLocalIPv4Address()(string, error) {
    address, err := net.InterfaceAddrs()
    if err != nil {
        return "", err
    }
    
    for _, a := range address {
        if ipnet, ok :=a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            return ipnet.IP.String(), nil
        }
    }

    return "", errors.New("Not found ipv4 address")
}

