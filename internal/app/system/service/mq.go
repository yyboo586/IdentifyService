package service

import (
	"context"
	"log"
	"sync"

	mqsdk "github.com/yyboo586/MQSDK"
)

type IMQ interface {
	Publish(ctx context.Context, topic string, userIDs []string, content map[string]interface{}) error
}

var localMQ IMQ

func MQ() IMQ {
	if localMQ == nil {
		localMQ = NewMQ()
	}
	return localMQ
}

func RegisterMQ(i IMQ) {
	localMQ = i
}

var (
	mqOnce     sync.Once
	mqInstance *mq
)

var config = &mqsdk.NSQConfig{
	Type:     "nsq",
	NSQDAddr: "124.221.243.128:4150",
	// 不使用NSQLookup，直接连接NSQD
	// NSQLookup: []string{},
}

type mq struct {
	producer mqsdk.Producer
}

func NewMQ() *mq {
	mqOnce.Do(func() {
		factory := mqsdk.NewFactory()

		// 创建生产者
		producer, err := factory.NewProducer(config)
		if err != nil {
			log.Fatalf("Failed to create producer: %v", err)
		}

		mqInstance = &mq{
			producer: producer,
		}
	})
	return mqInstance
}

func (m *mq) Publish(ctx context.Context, topic string, userIDs []string, content map[string]interface{}) error {
	body := map[string]interface{}{
		"user_ids": userIDs,
		"content":  content,
	}
	msg := &mqsdk.Message{
		Body: body,
	}
	return m.producer.Publish(ctx, topic, msg)
}
