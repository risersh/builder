package main

import (
	"log"

	"github.com/mateothegreat/go-rabbitmq/management"
	"github.com/risersh/builder/conf"
)

func main() {
	exchange := management.Exchange{
		Name:    "primary",
		Type:    "topic",
		Durable: true,
		Queues: []management.Queue{
			{
				Name:    "broker",
				Durable: true,
			},
		},
	}

	m := management.Management{}
	err := m.Connect(conf.Config.RabbitMQ.URI, management.SetupArgs{
		Exchanges: []management.Exchange{exchange},
	})
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
}
