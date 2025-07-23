package service

import (
	"IdentifyService/library/libWebsocket"
	"context"
	"errors"
	"log"

	"github.com/gogf/gf/v2/frame/g"
	mq "github.com/yyboo586/MQSDK"
)

type IMQ interface {
	Subscribe(ctx context.Context, topic string, channel string, handler mq.MessageHandler) error
}

var (
	localMQ *MQ
)

func MQService() IMQ {
	if localMQ == nil {
		panic("implement not found for interface IMQ, forgot register?")
	}
	return localMQ
}

const (
	TopicDeviceOnline  = "core.device.online"
	TopicDeviceOffline = "core.device.offline"
	TopicDeviceAlarm   = "core.device.alarm"
)

func RegisterMQService() {
	localMQ = NewMQ()
	localMQ.Subscribe(context.Background(), TopicDeviceOnline, "DeviceManagement1", localMQ.handleDeviceOnline)
	localMQ.Subscribe(context.Background(), TopicDeviceOffline, "DeviceManagement1", localMQ.handleDeviceOffline)
	localMQ.Subscribe(context.Background(), TopicDeviceAlarm, "DeviceManagement1", localMQ.handleDeviceAlarm)
}

type MQ struct {
	consumer mq.Consumer
}

func NewMQ() *MQ {
	nsqConfig := &mq.NSQConfig{
		Type:     "nsq",
		NSQDAddr: "124.220.236.38:4150",
		// 暂时不使用nsqlookupd，直接连接nsqd
		// NSQLookup: []string{"127.0.0.1:4160"},
	}
	factory := mq.NewFactory()
	consumer, err := factory.NewConsumer(nsqConfig)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	return &MQ{
		consumer: consumer,
	}
}

func (m *MQ) Subscribe(ctx context.Context, topic string, channel string, handler mq.MessageHandler) (err error) {
	err = m.consumer.Subscribe(ctx, topic, channel, handler)
	if err != nil {
		g.Log().Error(ctx, "subscribe failed", err)
	}
	return
}

func (m *MQ) handleDeviceOnline(msg *mq.Message) (err error) {
	ctx := context.Background()
	g.Log().Info(ctx, "device online", *msg)
	body, ok := msg.Body.(map[string]interface{})
	if !ok {
		g.Log().Error(ctx, "device online message body is not map[string]interface{}")
		return nil
	}

	orgID, ok := body["org_id"].(string)
	if !ok {
		g.Log().Error(ctx, "device online message body org_id is not string")
		return nil
	}

	data := map[string]interface{}{
		"device_id":   int64(body["device_id"].(float64)),
		"device_name": body["device_name"].(string),
		"device_key":  body["device_key"].(string),
	}

	MessagePush().PushToOrg(ctx, orgID, libWebsocket.EventDeviceOnline, data)
	return
}

func (m *MQ) handleDeviceOffline(msg *mq.Message) (err error) {
	ctx := context.Background()
	g.Log().Info(ctx, "device offline", *msg)
	body, ok := msg.Body.(map[string]interface{})
	if !ok {
		g.Log().Error(ctx, "device offline message body is not map[string]interface{}")
		return nil
	}

	orgID, ok := body["org_id"].(string)
	if !ok {
		g.Log().Error(ctx, "device offline message body org_id is not string")
		return nil
	}

	data := map[string]interface{}{
		"device_id":   int64(body["device_id"].(float64)),
		"device_name": body["device_name"].(string),
		"device_key":  body["device_key"].(string),
	}

	MessagePush().PushToOrg(ctx, orgID, libWebsocket.EventDeviceOffline, data)
	return
}

func (m *MQ) handleDeviceAlarm(msg *mq.Message) (err error) {
	g.Log().Info(context.Background(), "device alarm", *msg)

	var body map[string]interface{}
	body, ok := msg.Body.(map[string]interface{})
	if !ok {
		g.Log().Error(context.Background(), "device alarm message body is not map[string]interface{}")
		return errors.New("device alarm message body is not map[string]interface{}")
	}

	orgID, ok := body["org_id"].(string)
	if !ok {
		g.Log().Error(context.Background(), "device alarm message body org_id is not string")
		return errors.New("device alarm message body org_id is not string")
	}

	data := map[string]interface{}{
		"device_id":   int64(body["device_id"].(float64)),
		"device_name": body["device_name"].(string),
		"device_key":  body["device_key"].(string),
	}

	MessagePush().PushToOrg(context.Background(), orgID, libWebsocket.EventDeviceAlarm, data)

	return nil
}
