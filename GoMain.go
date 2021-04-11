package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	//http.HandleFunc("/", String("Hello, World!"))
	http.ListenAndServe(":8080", nil)
}
