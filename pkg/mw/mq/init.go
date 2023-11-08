package mq

import (
	"fmt"

	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"

	"github.com/streadway/amqp"
)

var RabbitMQ *amqp.Connection

func InitRabbitMQ() {
	connString := fmt.Sprintf("%s://%s:%s@%s:%s/",
		conf.GetConf().GetString("mq.RabbitMQ"),
		conf.GetConf().GetString("mq.RabbitMQUser"),
		conf.GetConf().GetString("mq.RabbitMQPassWord"),
		conf.GetConf().GetString("mq.RabbitMQHost"),
		conf.GetConf().GetString("mq.RabbitMQPort"),
	)
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	RabbitMQ = conn
}
