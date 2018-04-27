package main

import (
	//"fmt"
	l4g "base/log4go"
	"io/ioutil"
	"net/http"
)

//协议注册方法
func InitHttpCommand() {
	g_HttpCommandM.Register("log", UserLogCommand{})
}

type UserLogCommand struct {
}

func (this UserLogCommand) Execute(r *http.Request, p *HttpHandlerPool) bool {
	re, _ := ioutil.ReadAll(r.Body)

	//1.将数据反序列化--2.进行逻辑上的处理
	if loginfos := UnmarshalLog(re); loginfos == nil {
		l4g.Error("params is not log standard %s", string(re))
	} else {
		//将逻辑处理后的数据传输到kafka
		g_log_kafka.SendLogsToKafka(loginfos)
	}

	/*--Test--*/
	//fmt.Println(string(re))
	//		g_logm.Show()
	//fmt.Fprintln(w, "success receive!")
	defer r.Body.Close()
	return true
}
