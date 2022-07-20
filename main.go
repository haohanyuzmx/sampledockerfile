package main

import (
	"net"
)

func main() {
	_, err := net.Dial("tcp", "baidu.com")
	if err != nil {
		return
	}
}
