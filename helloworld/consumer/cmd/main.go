package main

import (
	"kafka-go-example/conf"
	"kafka-go-example/helloworld/consumer"
)

func main() {
	topic := conf.Test_Topic
	consumer.Consume(topic)
}
