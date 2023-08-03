package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 这里就可以打开文件，读取网页内容(css,html,js)
	str := "aaaaa"
	w.Write([]byte(str))
}

func f2(w http.ResponseWriter, r *http.Request) {
	// GET方法的参数都放在url里。
	// 获取url里的参数
	query := r.URL.Query()
	param1 := query.Get("a")
	param2 := query.Get("b")
	fmt.Println("a:", param1, "b:", param2)

	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("aaaa"))
}

// http server
func main() {
	http.HandleFunc("/mao/", f1)
	http.HandleFunc("/hello/", f2)
	// 一个函数建立连接并提供服务
	http.ListenAndServe("127.0.0.1:8080", nil)
}
