package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/olivere/elastic"
	"sync"
)

var (
	wg sync.WaitGroup
)

type hust struct {
	User string
	Age  int
}

func main() {
	consumer, err := sarama.NewConsumer([]string{"10.246.104.46:9092"}, nil)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetURL("http://10.246.104.46:9200"))
	if err != nil {
		fmt.Println("elastic.NewClient err ", err)
		return
	}

	partitionList, err := consumer.Partitions("dn_itemconsume")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(partitionList))

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("dn_itemconsume", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		wg.Add(1)
		defer pc.AsyncClose()

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

				//				body := map[string]interface{}{
				//					"first_name": "John",
				//					"last_name":  "Smith",
				//					"age":        25,
				//					"about":      "I love to go rock climbing",
				//				}

				var exist bool
				exist, err = client.IndexExists("tweets").Do(context.Background())
				fmt.Println("client.IndexExists exist ", exist)
				if err != nil {
					fmt.Println("client.IndexExists err ", err)
					continue
				}
				if !exist {
					_, err = client.CreateIndex("tweets").Do(context.Background())
					if err != nil {
						fmt.Println("client.CreateIndex err ", err)
						continue
					}
				}
				body := hust{"kk", 18}
				_, err = client.Index().Index("tweets").Type("doc").BodyJson(body).Do(context.Background())
				if err != nil {
					fmt.Println("client.Index err ", err)
					continue
				}
				fmt.Println("test++++")
			}
			fmt.Println("test---====")
		}(pc)
	}

	wg.Wait()
	consumer.Close()

}
