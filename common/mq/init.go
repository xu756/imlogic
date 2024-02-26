package mq

//
//type Connection struct {
//	*amqp.Connection
//}
//
//func Init() {
//	//声明连接对象
//	conn, err := amqp.Dial(config.RunData.MqUrl)
//	if err != nil {
//		log.Print(err)
//		return
//	}
//	defer conn.Close()
//	// 声明一个通道
//	ch, err := conn.Channel()
//	if err != nil {
//		log.Print(err)
//		return
//	}
//	defer ch.Close()
//	//创建队列
//	/*
//		参数1（name）：队列名字
//		参数2（durable）：持久化，队列中所有的数据都是在内存中的，如果为true的话，这个通道关闭之后，数据就会存在磁盘中持久化，false的话就会丢弃
//		参数3（autoDelete）：不需要用到队列的时候，是否将消息删除
//		参数4（exclusive）：是否独占队列，true的话，就是只能是这个进程独占这个队列，其他都不能对这个队列进行读写
//		参数5（noWait）：是否阻塞
//		参数6（args）：其他参数
//	*/
//	q, err := ch.QueueDeclare(
//		"hello", // 队列名字
//		false,   // 是否持久化，
//		false,   // 不用的时候是否自动删除
//		false,   // 用来指定是否独占队列
//		false,   // no-wait
//		nil,     // 其他参数
//	)
//	if err != nil {
//		log.Print(err)
//		return
//	}
//	log.Print(q)
//}
