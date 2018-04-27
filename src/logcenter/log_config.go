package main

import ()

type XmlConfig struct {
	Log     xmlLog      `xml:"log"`
	Server  xmlServer   `xml:"server"`
	Pprof   xmlPprof    `xml:"pprof"`
	Clients []xmlClient `xml:"client"`
	Amount  xmlAmount   `xml:"amount"`
}

type xmlLog struct {
	Config string `xml:"config"`
}
type xmlPprof struct {
	State string `xml:"state,attr"`
	Port  string `xml:"port"`
}

type xmlServer struct {
	Port string `xml:"port"`
}

type xmlClient struct {
	RemoteServer xmlRemoteServer `xml:"server"`
}

type xmlAmount struct {
	KafkaProducer uint32 `xml:"kafka_producer"`
	KafkaConsumer uint32 `xml:"kafka_consumer"`
	HandlePool    uint32 `xml:"handle_pool"`
}
type xmlRemoteServer struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

func (this *XmlConfig) RemoteServerAddr(name string) string {
	for _, v := range this.Clients {
		if v.RemoteServer.Name == name {
			return v.RemoteServer.Value
		}
	}
	return ""
}
