package main

/**
 * 文件名: logcenter.go
 * 创建时间: 2018年3月9日-下午4:47:10
 * 简介:
 * 详情: 游戏日志服
 */

import (
	"base/common"
	l4g "base/log4go"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

var (
	g_config       = new(XmlConfig)
	g_handler_chan = make(chan *HttpRequestInfo, 102400)
	g_kafka_addr   string
	g_logM         = new(LogM)
)

var configFile = flag.String("config", "config/logcenter.config.xml", "")

func init() {
	InitHttpCommand()
}

func main() {
	defer time.Sleep(time.Second * 3)
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	if err := common.LoadConfig(*configFile, g_config); err != nil {
		panic(fmt.Sprintf("load config %v fail: %v", *configFile, err))
	}
	l4g.LoadConfiguration(g_config.Log.Config)
	defer l4g.Close()

	g_kafka_addr = g_config.RemoteServerAddr("kafka")
	if g_kafka_addr == "" {
		l4g.Error("kafka config error")
		return
	}

	if g_config.Amount.HandlePool <= 0 || g_config.Amount.KafkaConsumer <= 0 || g_config.Amount.KafkaProducer <= 0 {
		l4g.Error("Amount config error")
		return
	}

	if g_config.Pprof.State == "true" {
		go func() {
			l4g.Error(http.ListenAndServe(g_config.Pprof.Port, nil))
		}()
	}

	//http pools
	for i := uint32(0); i < g_config.Amount.HandlePool; i++ {
		pool := NewHttpHandlerPool(int(i))
		go pool.Process()
	}

	//Kafka client
	g_log_kafka = NewLogKafka()
	if g_log_kafka == nil {
		l4g.Error("NewLogKafka error")
		return
	} else {
		l4g.Info("NewLogKafka success")
	}

	httpRouter()
	err := http.ListenAndServe(g_config.Server.Port, nil) //设置监听的端口
	if err != nil {
		l4g.Error("ListenAndServe: %v", err)
	}
}
