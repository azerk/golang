package main

import (
	"fmt"
	"net/http"
	//	"strings"
)

func getName(respone http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println(request.Form)
}
func main() {
	http.HandleFunc("/hi", getName)
}
