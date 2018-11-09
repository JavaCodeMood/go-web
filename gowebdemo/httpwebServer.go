package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm()   //解析参数，默认是不会解析的
	//在服务器端输出打印信息
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v :=range r.Form {
		fmt.Println("Key:",k)
		fmt.Println("Value:", strings.Join(v, ""))
	}
	//这个写入到w的是输出到客户端的
	fmt.Println(w, "Hello go web !")
}

func main(){
	//设置访问的路由
	http.HandleFunc("/",sayhelloName)
	//设置监听的端口
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe：", err)
	}
}
