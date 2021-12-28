package main

import (
	"kafka-go-example/conf"
	"kafka-go-example/helloworld/producer"
)

func main() {
	topic := conf.Test_Topic
	producer.Produce(topic, 1000)
}
