package main

import (
	l4g "base/log4go"
	"github.com/Shopify/sarama"
)

//var (
//	MAX_KAFKA_PRODUCER_POOL_SIZE int = 30 //kafka生产者连接池最大连接数
//	MAX_KAFKA_CONSUMER_POOL_SIZE int = 30 //kafka消费者连接池最大连接数
//)

func (this *LogKafka) PutProducer(producer sarama.AsyncProducer) {
	if uint32(len(this.producers)) >= g_config.Amount.KafkaProducer {
		producer.Close()
		return
	}
	this.producers <- producer
}

func (this *LogKafka) GetProducer() sarama.AsyncProducer {
	if len(this.producers) == 0 {
		producer, err := NewKafkaProducer()
		if err != nil {
			l4g.Error("sarama.NewAsyncProducer err: %v %s", err, g_kafka_addr)
			return nil
		}
		return producer
	}

	return <-this.producers
}

func (this *LogKafka) PutConsumer(consumer sarama.Consumer) {
	if uint32(len(this.consumers)) >= g_config.Amount.KafkaConsumer {
		consumer.Close()
		return
	}
	this.consumers <- consumer
}

func (this *LogKafka) GetConsumer() sarama.Consumer {
	if len(this.consumers) == 0 {
		consumer, err := sarama.NewConsumer([]string{g_kafka_addr}, nil)
		if err != nil {
			l4g.Error("sarama.NewConsumer error : %v %s", err, g_kafka_addr)
			return nil
		}
		return consumer
	}

	return <-this.consumers
}
