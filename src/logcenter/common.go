package main

/**
 * 文件名: common.go
 * 创建时间: 2018年3月9日-下午4:47:10
 * 简介:
 * 详情: 全局通用接口
 */

import ()

var (
	g_log_kafka *LogKafka
)

var g_topics []string = []string{KAFKA_LOGIN, KAFKA_REGISTER, KAFKA_OUTPUT, KAFKA_CONSUME, KAFKA_USERONLINE, KAFKA_STAGE, KAFKA_ONLINETIME, KAFKA_LEVEL, KAFKA_SILVERCONSUME}
