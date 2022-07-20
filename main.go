package main

import (
	"fmt"
	"net/http"
)

func main() {
	_, err := http.Get("http://baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
}
