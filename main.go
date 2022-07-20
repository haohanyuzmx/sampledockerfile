package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}
