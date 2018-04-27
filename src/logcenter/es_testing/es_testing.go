package main

import (
	l4g "base/log4go"
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"time"
)

var g_es_addr = "http://10.246.95.100:9200/"

func main() {
	client, err := elastic.NewClient(elastic.SetURL(g_es_addr))
	if err != nil {
		fmt.Printf("elastic.NewClient error : %v %s %v", err, g_es_addr, client)
		return
	}
	termQuery1 := elastic.NewTermQuery("userId", "920001000001")

	termQuery2 := elastic.NewRangeQuery("time").From("1523348015").To("1523348118")

	termQuery := elastic.NewBoolQuery().Must(termQuery1).Must(termQuery2)

	res, err := client.Search().Index("dn_login").Query(termQuery).Pretty(true).Do(context.Background())
	if err != nil {
		fmt.Printf("client.Index err %v", err)
		return
	}

	for _, v := range res.Hits.Hits {
		l4g.Info("SendKafkaLogToEs Index info %+v----- ", string(*v.Source))
	}
	time.Sleep(500 * time.Millisecond)
}
