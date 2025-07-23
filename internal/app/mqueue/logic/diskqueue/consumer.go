/**
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/12 10:22
 */

package diskqueue

import (
	"IdentifyService/internal/app/mqueue/model"
	"IdentifyService/internal/app/mqueue/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type consumer struct {
	Topic   string
	Channel string
	ID      string
	Handler model.MQConsumerHandlerCallback
}

// NewDiskConsumer 创建一个消费者
func NewDiskConsumer(topic, channel string, handler model.MQConsumerHandlerCallback) (service.IConsumer, error) {
	dqs, err := getClient(topic)
	if err != nil {
		return nil, err
	}
	id := "dqc_" + grand.S(16)
	c := &consumer{
		Topic:   topic,
		Channel: channel,
		ID:      id,
		Handler: handler,
	}
	dqs.RegisterConsumer(channel, c)
	return c, nil
}

func (s *consumer) CloseMqConsumer() {
	dqs, err := getClient(s.Topic)
	if err != nil {
		g.Log("diskQueue").Error(context.TODO(), "执行 CloseMqConsumer 失败："+err.Error())
		return
	}
	dqs.RemoveConsumer(s.Channel, s.ID)
}
