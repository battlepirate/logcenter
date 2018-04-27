package main

import (
	l4g "base/log4go"
	"encoding/json"
	"github.com/Shopify/sarama"
	"time"
)

type LogKafka struct {
	//producer  sarama.AsyncProducer
	producers chan sarama.AsyncProducer
	//consumers map[string]sarama.Consumer
	consumers chan sarama.Consumer
}

func NewKafkaProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	//config.Producer.Return.Successes = true
	config.Producer.Flush.Frequency = 500 * time.Millisecond
	producer, err := sarama.NewAsyncProducer([]string{g_kafka_addr}, config)
	return producer, err
}
func NewLogKafka() *LogKafka {
	result := &LogKafka{}
	result.producers = make(chan sarama.AsyncProducer, g_config.Amount.KafkaProducer)
	result.consumers = make(chan sarama.Consumer, g_config.Amount.KafkaConsumer)

	for i := uint32(0); i < g_config.Amount.KafkaProducer; i++ {
		producer, err := NewKafkaProducer()
		if err != nil {
			l4g.Error("NewKafkaProducer err: %v %s", err, g_kafka_addr)
			return nil
		}
		result.producers <- producer
		l4g.Info("New Kafka Producer pool success: %d ", i)
	}
	//consumers := make(map[string]sarama.Consumer)

	for i := uint32(0); i < g_config.Amount.KafkaConsumer; i++ {
		consumer, err := sarama.NewConsumer([]string{g_kafka_addr}, nil)
		if err != nil {
			l4g.Error("sarama.NewConsumer error : %v %s", err, g_kafka_addr)
			return nil
		}
		result.consumers <- consumer
		l4g.Info("New Kafka consumer pool success: %d ", i)
	}

	//result.producer = producer

	return result

}

func (this *LogKafka) SendLogsToKafka(infos []LogInfo) bool {
	for _, info := range infos {
		switch info.MainType {
		case LOG_MAIN_SERVER_INFO: //服务器变化
			switch info.ChildType {
			case LOG_CHILD_SERVER_PCU: //在线人数
				data, _ := json.Marshal(ChangeLogToUserOnline(info))
				this.SendLogToKafka(data, KAFKA_USERONLINE)
			}

		case LOG_MAIN_USER_CHANGE: //玩家变化
			switch info.ChildType {
			case LOG_CHILD_USER_LOGIN: //登录
				data, _ := json.Marshal(ChangeLogToLogin(info))
				this.SendLogToKafka(data, KAFKA_LOGIN)
			case LOG_CHILD_USER_CREATE: //创角
				data, _ := json.Marshal(ChangeLogToRegister(info))
				this.SendLogToKafka(data, KAFKA_REGISTER)
			case LOG_CHILD_USER_OFFLINE: //在线时长
				data, _ := json.Marshal(ChangeLogToOnlineTime(info))
				this.SendLogToKafka(data, KAFKA_ONLINETIME)
			case LOG_CHILD_USER_UPGRADE: //升级
				data, _ := json.Marshal(ChangeLogToLevel(info))
				this.SendLogToKafka(data, KAFKA_LEVEL)
			}

		case LOG_MAIN_RESOURCE_CHANGE: //资源变化
			switch info.ChildType {
			case LOG_CHILD_RESOURCE_ADD: //资源产出
				data, _ := json.Marshal(ChangeLogToOutPut(info))
				this.SendLogToKafka(data, KAFKA_OUTPUT)
			case LOG_CHILD_RESOURCE_DEC: //资源消耗
				data, _ := json.Marshal(ChangeLogToConsume(info))
				this.SendLogToKafka(data, KAFKA_CONSUME)
			case LOG_CHILD_RESOURCE_DIAMOND_ADD, LOG_CHILD_RESOURCE_DIAMOND_DEC: //钻石获得 ,钻石消耗
				data, _ := json.Marshal(ChangeLogToSilverConsume(info))
				this.SendLogToKafka(data, KAFKA_SILVERCONSUME)
			}
		case LOG_MAIN_ACTION: //玩家行为
			switch info.ChildType {
			case LOG_CHILD_ACTION_STAGE: //关卡
				data, _ := json.Marshal(ChangeLogToStage(info))
				this.SendLogToKafka(data, KAFKA_STAGE)

			}
		default:
			return true
		}
	}
	return true
}

func (this *LogKafka) SendLogToKafka(info interface{}, topics string) bool {
	data, ok := info.([]byte)
	if !ok {
		l4g.Error("interface not  []byte: %v", info)
		return false
	}
	g_log_kafka.asyncProducer(data, topics)
	return true
}

//异步生产者:适合高并发
func (this *LogKafka) asyncProducer(data []byte, topics string) bool {
	producer := g_log_kafka.GetProducer()
	if producer == nil {
		l4g.Error("g_log_kafka.GetProducer() Errors:  %s", topics)
		return false
	}

	go func() {
		for err := range producer.Errors() {
			l4g.Error("g_log_kafka.producer.Errors: %v %s", err, topics)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: topics,
	}

	msg.Value = sarama.ByteEncoder(data)
	producer.Input() <- msg

	g_log_kafka.PutProducer(producer)

	//g_logM.InTest()
	return true
}

//同步生产者:小并发量,大并发量会失去响应,不适用
func (this *LogKafka) syncProducer(data []byte, topics string) bool {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer([]string{g_kafka_addr}, config)
	if err != nil {
		l4g.Error("arama.NewSyncProducer err: %v %s", err, topics)
		return false
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topics,
	}

	msg.Value = sarama.ByteEncoder(data)
	l4g.Info("kafka value is ---- %v", msg.Value)

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		l4g.Error("Kafka Send message Fail: %v", err)
		return false
	}
	l4g.Info("kafka result is ----topics =%s Partition = %d, offset=%d", topics, partition, offset)
	return true
}
