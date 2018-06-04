package main

import (
	"net/http"
	"fmt"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "hello world")
}

func main()  {
	http.HandleFunc("/",IndexHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
	
}