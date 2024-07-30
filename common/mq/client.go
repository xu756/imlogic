package mq

import (
	"imlogic/common/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

// rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl string
}

// 创建结构体实例
func newRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: config.RunData.MqUrl}
}

// 断开channel 和 connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}
