package main

import (
	l4g "base/log4go"
	//"fmt"
	"net/http"
)

func httpRouter() {
	http.HandleFunc("/", RouterHandler)
}
func RouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form.Get("action") == "" {
		l4g.Error("action is nil")
		return
	}
	if r.Method == "POST" {
		closeChan := make(chan bool, 0)
		g_handler_chan <- NewHttpRequestInfo(r.Form.Get("action"), r, closeChan)
		_ = <-closeChan
	} else {
		l4g.Error("Method is not POST")
		//fmt.Fprintln(w, "Method is not POST!")
	}
}
