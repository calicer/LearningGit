package main

import (
	"fmt"
	sarama "github.com/bsm/sarama-cluster"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	//config.Consumer.Return.Notifications = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.CommitInterval = time.Second

	brokers := []string{"localhost:9092"}

	consumer, err := sarama.NewConsumer(brokers, "test-group", []string{"TestTopic"}, config)
	channel := consumer.Messages()

	fmt.Println(err)
	fmt.Println(&channel)

	go func() {
		fmt.Println("++++++++++++++++++++")

		for err := range consumer.Errors() {
			fmt.Println(err)
			fmt.Println("++++++++++++++++++++")
		}
	}()

	// go func(){
	// 	for ntf := range consumer.Notifications(){
	// 		fmt.Println(ntf)
	// 	}
	// }

	for {
		select {
		case msg := <-channel:
			fmt.Println(string(msg.Value))
		}

	}

}
