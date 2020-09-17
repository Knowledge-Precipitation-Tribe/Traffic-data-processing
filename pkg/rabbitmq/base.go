package rabbitmq

import (
	"Traffic-data-processing/config"
	"Traffic-data-processing/pkg/logging"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-09-16 21:00
* @Description:
**/

var logger = logging.GetLogger()

type RabbitMQ struct {
	conn    *amqp.Connection
	Channel *amqp.Channel

	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	Mqurl string
}

// 创建RabbitMQ实例
func NewRabbitMQ(queuqName string,
	exchange string, key string) *RabbitMQ {

	rabbitMQUrl, err := config.GetRabbitMQUrl()
	if err != nil{
		logger.Error("get rabbitmq url error", zap.Error(err))
		panic(err)
	}
	rabbitmq := &RabbitMQ{
		QueueName: queuqName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     rabbitMQUrl,
	}
	//创建rabbitmq连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailOnErr(err, "创建连接错误")
	rabbitmq.Channel, err = rabbitmq.conn.Channel()
	rabbitmq.FailOnErr(err, "获取channel失败")
	return rabbitmq
}

// 断开channel和connection的连接释放资源
func (r *RabbitMQ) Destroy() {
	r.Channel.Close()
	r.conn.Close()
}

//自定义错误处理函数
func (r *RabbitMQ) FailOnErr(err error, message string) {
	if err != nil {
		logging.GetLogger().Error(message, zap.Error(err))
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}