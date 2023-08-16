package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("./lesson02/hello.tmpl")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	err02 := templ.Execute(w, "文字")
	if err02 != nil {
		fmt.Printf(err02.Error())
		return
	}
}

// template
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err: %v\n", err)
		return
	}
}
