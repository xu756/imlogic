package mq

import (
	"imlogic/common/config"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 定义客户端结构体
type Client struct {
	conn      *amqp.Connection // AMQP 连接
	channel   *amqp.Channel    // AMQP 通道
	QueueName string           // 队列名称
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
	// 声明交换机
	err = channel.ExchangeDeclare(
		"chat",   // 交换机名称
		"direct", // 交换机类型（路由）
		true,     // 是否持久化
		false,    // 是否自动删除
		false,    // 是否内部使用
		false,    // 是否立即返回
		nil,      // 其他参数
	)
	if err != nil {
		panic(err) // 如果声明交换机失败，程序崩溃
	}

	// 声明群聊交换机
	err = channel.ExchangeDeclare(
		"chat_group", // 群聊交换机名称
		"fanout",     // 交换机类型（扇形广播）
		true,         // 是否持久化
		false,        // 是否自动删除
		false,        // 是否内部使用
		false,        // 是否立即返回
		nil,          // 其他参数
	)
	if err != nil {
		panic(err) // 如果声明交换机失败，程序崩溃
	}

	return &Client{
		conn:      conn,
		channel:   channel,
		QueueName: queueName,
	}
}

// 发布消息到交换机
func (c *Client) WorkPublish(message, messageType string) error {
	var exchange string
	var routingKey string

	if messageType == "group" {
		exchange = "chat_group"
		routingKey = "" // 对于群聊，不需要路由键
	} else {
		exchange = "chat"
		routingKey = c.QueueName
	}

	err := c.channel.Publish(
		exchange,   // 交换机名称
		routingKey, // 路由键
		false,      // 如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,      // 如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		amqp.Publishing{
			ContentType: "text/plain",    // 内容类型
			Body:        []byte(message), // 消息正文
		})
	if err != nil {
		return err // 如果发布消息失败，返回错误
	}

	return nil // 成功返回
}

// 消费者
func (c *Client) Consume(callback func(delivery amqp.Delivery)) error {
	// 1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := c.channel.QueueDeclare(
		c.QueueName,
		true,  // 是否持久化
		false, // 是否自动删除
		false, // 是否具有排他性
		false, // 是否阻塞处理
		nil,   // 额外的属性
	)
	if err != nil {
		return err
	}

	// 绑定队列到群聊交换机
	err = c.channel.QueueBind(
		q.Name,       // 队列名称
		"",           // 路由键（对于群聊交换机无效）
		"chat_group", // 交换机名称
		false,        // 不等待
		nil,          // 额外的参数
	)
	if err != nil {
		return err
	}

	// 接收消息
	msgs, err := c.channel.Consume(
		q.Name, // queue
		"",     // 用来区分多个消费者
		false,  // 是否自动应答
		false,  // 是否独有
		false,  // 设置为true，表示 不能将同一个Connection中生产者发送的消息传递给这个Connection中的消费者
		false,  // 列是否阻塞
		nil,    // 额外的参数
	)
	if err != nil {
		return err
	}

	go func() { // 启动一个goroutine来处理消息
		for d := range msgs { // 遍历接收到的消息
			callback(d)  // 调用回调函数处理消息
			d.Ack(false) // 确认消息已处理
		}
	}()

	return nil // 成功返回
}

// 关闭连接
func (c *Client) Close() {
	c.channel.Close() // 关闭通道
	c.conn.Close()    // 关闭连接
}
