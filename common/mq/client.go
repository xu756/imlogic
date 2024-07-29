package mq

import (
	"imlogic/common/config"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 定义客户端结构体
type Client struct {
	conn    *amqp.Connection // AMQP 连接
	channel *amqp.Channel    // AMQP 通道
	// 队列名称
	QueueName string
	// 交换机名称
	Exchange string
}

// 创建一个新的客户端实例
func NewClient(queueName string) *Client {
	conn, err := amqp.Dial(config.RunData.MqUrl) // 建立到RabbitMQ服务器的连接
	if err != nil {
		panic(err) // 如果连接失败，程序崩溃
	}
	log.Print("mq connected")
	channel, err := conn.Channel() // 打开一个新的通道
	if err != nil {
		panic(err) // 如果打开通道失败，程序崩溃
	}
	log.Print("mq channel opened")
	return &Client{ // 返回新的客户端实例
		conn:      conn,
		channel:   channel,
		QueueName: queueName,
		Exchange:  "imlogic", // 默认交换机名称
	}
}

// 发布消息到交换机
func (c *Client) Publish(message string) error {
	err := c.channel.ExchangeDeclare( // 声明交换机
		c.Exchange, // 交换机名称
		"fanout",   // 交换机类型（扇形广播）
		true,       // 是否持久化
		false,      // 是否自动删除
		false,      // 是否内部使用
		false,      // 是否立即返回
		nil,        // 其他参数
	)
	if err != nil {
		return err // 如果声明交换机失败，返回错误
	}

	err = c.channel.Publish( // 发布消息
		c.Exchange, // 交换机名称
		"",         // 路由键（对于扇形广播交换机，此字段为空）
		false,      // 是否启用mandatory标志
		false,      // 是否启用immediate标志
		amqp.Publishing{ // 消息属性
			ContentType: "text/plain",    // 内容类型
			Body:        []byte(message), // 消息正文
		})
	if err != nil {
		return err // 如果发布消息失败，返回错误
	}

	return nil // 成功返回
}

// 订阅消息并处理
func (c *Client) Consume(callback func(delivery amqp.Delivery)) error {
	err := c.channel.ExchangeDeclare( // 声明交换机
		c.Exchange, // 交换机名称
		"fanout",   // 交换机类型（扇形广播）
		true,       // 是否持久化
		false,      // 是否自动删除
		false,      // 是否内部使用
		false,      // 是否立即返回
		nil,        // 其他参数
	)
	if err != nil {
		return err // 如果声明交换机失败，返回错误
	}

	_, err = c.channel.QueueDeclare( // 声明队列
		c.QueueName, // 队列名称
		true,        // 是否持久化
		false,       // 是否自动删除
		false,       // 是否排他性
		false,       // 是否立即返回
		nil,         // 其他参数
	)
	if err != nil {
		return err // 如果声明队列失败，返回错误
	}

	err = c.channel.QueueBind( // 绑定队列到交换机
		c.QueueName, // 队列名称
		"",          // 路由键（对于扇形广播交换机，此字段为空）
		c.Exchange,  // 交换机名称
		false,       // 是否立即返回
		nil,         // 其他参数
	)
	if err != nil {
		return err // 如果绑定队列失败，返回错误
	}

	msgs, err := c.channel.Consume( // 开始消费消息
		c.QueueName, // 队列名称
		"",          // 消费者标签
		false,       // 是否自动确认
		false,       // 是否排他性
		false,       // 是否只从本地节点消费
		false,       // 是否立即返回
		nil,         // 其他参数
	)
	if err != nil {
		return err // 如果开始消费失败，返回错误
	}

	go func() { // 启动一个goroutine来处理消息
		for d := range msgs { // 遍历接收到的消息
			callback(d)  // 调用回调函数处理消息
			d.Ack(false) // 确认消息已处理
		}
	}()

	return nil // 成功返回
}
