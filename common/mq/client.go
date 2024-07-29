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
func NewClient(queueName, exchange string) *Client {
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

	// err = channel.ExchangeDeclare( // 声明交换机
	// 	"exchange", // 交换机名称
	// 	"fanout",   // 交换机类型（扇形广播）
	// 	true,       // 是否持久化
	// 	false,      // 是否自动删除
	// 	false,      // 是否内部使用
	// 	false,      // 是否立即返回
	// 	nil,        // 其他参数
	// )
	// if err != nil {
	// 	panic(err) // 如果声明交换机失败，程序崩溃
	// }
	return &Client{ // 返回新的客户端实例
		conn:      conn,
		channel:   channel,
		QueueName: queueName,
		Exchange:  exchange, // 默认交换机名称
	}
}

// 发布消息到交换机
func (c *Client) WorkPublish(message string) error {
	_, err := c.channel.QueueDeclare( // 声明交换机
		c.QueueName,
		//是否持久化
		true,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		return err // 如果声明交换机失败，返回错误
	}

	err = c.channel.Publish( // 发布消息
		c.Exchange, // 交换机名称
		c.QueueName,
		false, //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false, //如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		amqp.Publishing{ // 消息属性
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
		//是否持久化
		true,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		return err
	}

	// 接收消息
	msgs, err := c.channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"",
		//是否自动应答
		false, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
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
