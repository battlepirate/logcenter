package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer([]string{"10.241.104.48:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     "scene_log",
		Partition: int32(1),
		Key:       sarama.StringEncoder("key"),
	}
	var value string
	var data = `{"name":"Xiao mi 6","product_id":"10","number":"10000","price":"2499","is_on_sale":"true"}`
	for {
		_, err := fmt.Scanf("%v", &value)
		if err != nil {
			fmt.Println("fmt.Scanf:", err, value)
			continue
		}
		msg.Key = sarama.ByteEncoder([]byte("wc_test_key"))
		msg.Value = sarama.ByteEncoder(data)
		fmt.Println(value)
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("Send message Fail:", err)
		}
		fmt.Println("Partition = %d, offset=%d\n", partition, offset)
	}
}
