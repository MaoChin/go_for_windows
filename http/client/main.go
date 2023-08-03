package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http client
func main() {
	// url中 ? 后为参数
	res, err := http.Get("http://127.0.0.1:8080/hello?a=1&b=111")
	if err != nil {
		fmt.Printf("Get error:%v\n", err)
		return
	}
	ret, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("ReadAll error:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}
