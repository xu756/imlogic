package mq

import (
	"encoding/json"
	"imlogic/common/types"
	"imlogic/kitex_gen/im"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 私聊MQ
func NewPrivateMessageMQ() (rabbitmq *RabbitMQ, err error) {
	// 创建RabbitMQ实例
	rabbitmq = newRabbitMQ("private", "", "")
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	if err != nil {
		return nil, err
	}
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		return nil, err
	}
	return rabbitmq, nil
}

// 发布私聊消息
func (r *RabbitMQ) PublishPrivateMessage(linkId, HostName string, msg *im.Message) error {
	// 1. 申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	// 2. 序列化消息
	data, err := json.Marshal(&types.MqPrivateMessage{
		Msg:      msg,
		HostName: HostName,
		LinkId:   linkId,
	})
	if err != nil {
		return err
	}
	// 2. 发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// 消费私聊消息
func (r *RabbitMQ) ConsumePrivateMessage(f func(linkId, HostName string, msg *im.Message)) (err error) {
	// 1. 申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err = r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	// 2. 消费消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	// 3. 处理消息
	go func() {
		for msg := range msgs {
			m := new(types.MqPrivateMessage)
			err := json.Unmarshal(msg.Body, m)
			if err != nil {
				continue
			}
			f(m.LinkId, m.HostName, m.Msg)
		}
	}()
	make(chan bool) <- true
	return nil
}

// 广播（全体）MQ
func NewBroadcastMessageMQ() (rabbitmq *RabbitMQ, err error) {
	// 创建RabbitMQ实例
	rabbitmq = newRabbitMQ("", "broadcast", "")
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	if err != nil {
		return nil, err
	}
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		return nil, err
	}
	return rabbitmq, nil
}

// 发布广播消息
func (r *RabbitMQ) PublishBroadcastMessage(msg *im.Message) (err error) {
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = r.channel.Publish(
		r.Exchange,
		"fanout",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// 消费广播消息
func (r *RabbitMQ) ConsumeBroadcastMessage(f func(msg *im.Message)) (err error) {
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"",
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	err = r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	go func() {
		for msg := range msgs {
			m := new(im.Message)
			err := json.Unmarshal(msg.Body, m)
			if err != nil {
				continue
			}
			f(m)
		}
	}()
	make(chan bool) <- true
	return nil
}

// 延迟任务 3分钟后执行
func NewDelayMessageMQ(name string) (rabbitmq *RabbitMQ, err error) {
	// 创建RabbitMQ实例
	rabbitmq = newRabbitMQ("", name, "delay")
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	if err != nil {
		return nil, err
	}
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		return nil, err
	}
	return rabbitmq, nil
}

// 发布延迟消息
func (r *RabbitMQ) PublishDelayMessage(delaySecend int, message string) (err error) {
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		"x-delayed-message",
		true,
		false,
		false,
		false,
		amqp.Table{"x-delayed-type": "direct"},
	)
	if err != nil {
		return err
	}
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			Headers:     amqp.Table{"x-delay": delaySecend * 1000},
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// 消费延迟消息
func (r *RabbitMQ) ConsumeDelayMessage() (<-chan amqp.Delivery, error) {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"x-delayed-message",
		true,
		false,
		false,
		false,
		amqp.Table{"x-delayed-type": "direct"},
	)
	if err != nil {
		return nil, err
	}
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"",
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	err = r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}
