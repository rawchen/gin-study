package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("./lesson03/hello.tmpl")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	u1 := User{
		Name:   "名称",
		Gender: "男", // 测试小写开头
		Age:    12,
	}
	// gender is an unexported field of struct type main.Usertemplate:

	m1 := map[string]interface{}{
		"name":   "小王子",
		"gender": "男",
		"age":    19,
	}

	hobbyList := []string{
		"足球",
		"篮球",
		"羽毛球",
	}

	err02 := templ.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
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
