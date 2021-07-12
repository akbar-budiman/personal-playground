package main

import (
	"encoding/json"
	"fmt"

	"github.com/nsqio/go-nsq"
)

var (
	nsqAddress        = "127.0.0.1:4150"
	nsqlookupdAddress = "127.0.0.1:4161"
	addRectTopic      = "NewRectInserted"
	addCircleTopic    = "NewCircleInserted"
)

func ProduceNewRect(rect *Rect) {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(nsqAddress, config)
	if err != nil {
		panic(err)
	}

	newObj, err := json.Marshal(rect)
	if err != nil {
		panic(err)
	}
	messageBody := []byte(newObj)

	err = p.Publish(addRectTopic, messageBody)
	if err != nil {
		panic(err)
	}

	p.Stop()
}

func ProduceNewCircle(circle *Circle) {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(nsqAddress, config)
	if err != nil {
		panic(err)
	}

	newObj, err := json.Marshal(circle)
	if err != nil {
		panic(err)
	}
	messageBody := []byte(newObj)

	err = p.Publish(addCircleTopic, messageBody)
	if err != nil {
		panic(err)
	}

	p.Stop()
}

type ConsumeNewRectToInsert struct{}

func (h *ConsumeNewRectToInsert) HandleMessage(m *nsq.Message) error {
	var newObj Rect
	json.Unmarshal(m.Body, &newObj)

	SetData(newObj.Id, string(m.Body[:]))
	return nil
}

type ConsumeNewRectToReport struct{}

func (h *ConsumeNewRectToReport) HandleMessage(m *nsq.Message) error {
	fmt.Println("Got a new rect: ", m)
	return nil
}

func RegisterConsumer() {
	fmt.Println("Registering consumer")
	config := nsq.NewConfig()

	fmt.Println("Registering consumer : insertingConsumer")
	insertingConsumer, err := nsq.NewConsumer(addRectTopic, "channel1", config)
	if err != nil {
		panic(err)
	}
	insertingConsumer.AddHandler(&ConsumeNewRectToInsert{})

	err = insertingConsumer.ConnectToNSQLookupd(nsqlookupdAddress)
	if err != nil {
		panic(err)
	}

	fmt.Println("Registering consumer : reportingConsumer")
	reportingConsumer, err := nsq.NewConsumer(addRectTopic, "channel2", config)
	if err != nil {
		panic(err)
	}
	reportingConsumer.AddHandler(&ConsumeNewRectToReport{})

	err = reportingConsumer.ConnectToNSQLookupd(nsqlookupdAddress)
	if err != nil {
		panic(err)
	}

	fmt.Println("consumer registered.")
}
