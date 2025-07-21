package internal

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	kfkConsumer *kafka.Consumer
	processor   Processor
}

func NewConsumer(processor Processor, cfgMap *kafka.ConfigMap, topics ...string) (*Consumer, error) {
	c, err := kafka.NewConsumer(cfgMap)
	if err != nil {
		return nil, err
	}
	if err = c.SubscribeTopics(topics, nil); err != nil {
		return nil, err
	}

	return &Consumer{processor: processor, kfkConsumer: c}, nil
}

func (c *Consumer) Consume() error {
	for {
		ev := c.kfkConsumer.Poll(100)
		if ev == nil {
			continue
		}
		switch e := ev.(type) {
		case *kafka.Message:
			if err := c.processor.Process(e); err != nil {
				return err
			}
		case kafka.Error:
			fmt.Printf("Error [%s]\n", e.Error())
			return e
		case kafka.PartitionEOF:
			fmt.Printf("Reached end of partition")
		case kafka.OffsetsCommitted:
			continue
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
}
