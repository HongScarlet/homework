package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析參數，預設是不會解析的
	fmt.Println("======== For / ========")
	fmt.Println(r.Form)  //這些資訊是輸出到伺服器端的列印資訊
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("======== For / with parameter ========")
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //這個寫入到 w 的是輸出到客戶端的
}
func sayhelloName1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======== For /api1 ========")
	r.ParseForm()  //解析參數，預設是不會解析的
	fmt.Println(r.Form)  //這些資訊是輸出到伺服器端的列印資訊
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("======== For /api1 with parameter ========")
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!+/api1") //這個寫入到 w 的是輸出到客戶端的
}
func sayhelloName2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======== For /api1/ ========")
	r.ParseForm()  //解析參數，預設是不會解析的
	fmt.Println(r.Form)  //這些資訊是輸出到伺服器端的列印資訊
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("======== For /api1/ with parameter ========")
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!+/api1/") //這個寫入到 w 的是輸出到客戶端的
}

func main() {
	http.HandleFunc("/", sayhelloName) //設定訪問的路由
	http.HandleFunc("/api1", sayhelloName1) //設定訪問的路由	
	http.HandleFunc("/api1/", sayhelloName2) //設定訪問的路由	

	err := http.ListenAndServe(":9090", nil) //設定監聽的埠
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
































