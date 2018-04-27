package main

import (
	l4g "base/log4go"
	"net/http"
)

type HttpRequestInfo struct {
	action    string
	req       *http.Request
	closeChan chan bool
}

func NewHttpRequestInfo(action string, req *http.Request, c chan bool) *HttpRequestInfo {
	return &HttpRequestInfo{
		action:    action,
		req:       req,
		closeChan: c,
	}
}

type HttpHandlerPool struct {
	pool_id int
}

func NewHttpHandlerPool(i int) *HttpHandlerPool {
	return &HttpHandlerPool{
		pool_id: i,
	}
}
func (this *HttpHandlerPool) Process() {
	for {
		l4g.Debug("HttpHandlerPool processs %d", this.pool_id)
		select {
		case request := <-g_handler_chan:
			g_HttpCommandM.Dispatcher(request.action, request.req, this)
			request.closeChan <- true
		}
	}
	l4g.Debug("HttpHandlerPool processs %d finish", this.pool_id)
}
