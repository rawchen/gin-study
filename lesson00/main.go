package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fileByte, err := os.ReadFile("./lesson00/hello.txt")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	//_, err := fmt.Fprintln(w, "Hello Golang!")
	_, err02 := fmt.Fprintln(w, string(fileByte))
	if err02 != nil {
		fmt.Printf(err02.Error())
		return
	}
}

// web本质
func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err: %v\n", err)
		return
	}
}
