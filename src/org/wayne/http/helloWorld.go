package main

import (
	"net/http"
	"fmt"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "hello world")
}
// Clinet -> Requests ->  [Multiplexer(router) -> handler  -> Response -> Clinet
func main()  {
	http.HandleFunc("/",IndexHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
	
}