package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {

	// 自定义一个模板函数
	k := func(arg string) (string, error) {
		return arg + "测试！", nil
	}
	// 采用链式操作在Parse之前调用Funcs添加自定义的kua函数
	t := template.New("f.tmpl")

	t.Funcs(template.FuncMap{
		"key": k,
	})
	_, err := t.ParseFiles("./lesson04/f.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}

	name := "名称"

	// 使用name渲染模板，并将结果写入
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err: %v\n", err)
		return
	}

	// 结果
	//名称测试！
}
