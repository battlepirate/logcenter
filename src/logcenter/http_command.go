package main

import (
	l4g "base/log4go"
	"net/http"
)

var g_HttpCommandM = NewHttpCommandM()

type HttpCommand interface {
	Execute(*http.Request, *HttpHandlerPool) bool
}

type HttpCommandM struct {
	cmdm map[string]HttpCommand
}

func NewHttpCommandM() *HttpCommandM {
	return &HttpCommandM{
		cmdm: make(map[string]HttpCommand),
	}
}

func (this *HttpCommandM) Register(name string, cmd HttpCommand) {
	this.cmdm[name] = cmd
}

func (this *HttpCommandM) Dispatcher(action string, req *http.Request, pool *HttpHandlerPool) bool {
	l4g.Debug("[Command] Dispatcher cmd: %s", action)
	if cmd, exist := this.cmdm[action]; exist {
		cmd.Execute(req, pool)
		return true
	}
	l4g.Error("[Command] no find action cmd: %s", action)
	return false
}
