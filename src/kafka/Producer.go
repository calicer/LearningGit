package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {

	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_2
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	// producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)

	// fmt.Println("----")
	// //sarama.
	// //log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// producer.SendMessage(&sarama.ProducerMessage{Topic: "BrandTopic", Value: sarama.StringEncoder("Test message from code!!")})

	config.Consumer.Return.Errors = true
	//kafka end point
	brokers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Println(err)
	}

	str := Marshal()

	for i := 0; i <= 10; i++ {
		producer.SendMessage(&sarama.ProducerMessage{Topic: "TestTopic", Value: sarama.StringEncoder(str)})
		time.Sleep(4 * time.Second)
	}
	//get broker
	// cluster, err := sarama.NewConsumer(brokers, config)
	// if err != nil {
	// 	panic(err)
	// }

	// defer func() {
	// 	if err := cluster.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// //get all topic from cluster
	// topics, _ := cluster.Topics()
	// for index := range topics {
	// 	fmt.Println(topics[index])
	// }
}

type Employee struct {
	ID   string
	Name string
}

func Marshal() string {

	e := Employee{ID: "1", Name: "Vishal"}

	str, _ := json.Marshal(e)
	return string(str)
}
