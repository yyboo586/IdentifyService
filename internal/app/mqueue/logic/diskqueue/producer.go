/**
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/12 10:23
 */

package diskqueue

import (
	"IdentifyService/internal/app/mqueue/consts"
	"IdentifyService/internal/app/mqueue/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type producer struct {
	isClose bool
}

func NewDiskQueueProducer() service.IProducer {
	return &producer{}
}

func (p *producer) Publish(topic string, body []byte) error {
	dq, err := getClient(topic)
	if err != nil {
		return err
	}
	return dq.Publish(body)
}

func (p *producer) PublishAsync(topic string, body []byte) error {
	dq, err := getClient(topic)
	go func() {
		err = dq.Publish(body)
		if err != nil {
			g.Log("diskQueue").Error(context.TODO(), "diskQueue PublishAsync消息失败："+err.Error())
		}
	}()
	return err
}

func (p *producer) DelayPublish(topic string, body []byte, delay consts.MsgDelayLevel) error {
	g.Log("diskQueue").Warning(context.TODO(), "diskQueue 不支持延时消息，将使用publish发送")
	dq, err := getClient(topic)
	if err != nil {
		return err
	}
	return dq.Publish(body)
}

func (p *producer) Close() {
	p.isClose = true
}
