package service

import (
	"IdentifyService/library/libWebsocket"
	"context"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	messagePushOnce     sync.Once
	messagePushInstance IMessagePush
)

func MessagePush() IMessagePush {
	messagePushOnce.Do(func() {
		messagePushInstance = &messagePush{}
	})
	return messagePushInstance
}

type IMessagePush interface {
	// 基础推送方法
	PushToAll(ctx context.Context, event libWebsocket.WebSocketEvent, data interface{}) error
	PushToUser(ctx context.Context, userID string, event libWebsocket.WebSocketEvent, data interface{}) error
	PushToUsers(ctx context.Context, userIDs []string, event libWebsocket.WebSocketEvent, data interface{}) error
	PushToClient(ctx context.Context, clientID string, event libWebsocket.WebSocketEvent, data interface{}) error
	PushToOrg(ctx context.Context, orgID string, event libWebsocket.WebSocketEvent, data interface{}) error
	// 组织相关推送
	PushOrgCreated(ctx context.Context, orgID string, orgInfo interface{}) error
	PushOrgUpdated(ctx context.Context, orgID string, orgInfo interface{}) error
	PushOrgDeleted(ctx context.Context, orgID string) error
	PushOrgStatusChanged(ctx context.Context, orgID string, enabled bool) error

	// 用户相关推送
	PushUserCreated(ctx context.Context, userID string, userInfo interface{}) error
	PushUserUpdated(ctx context.Context, userID string, userInfo interface{}) error
	PushUserDeleted(ctx context.Context, userID string) error
	PushUserStatusChanged(ctx context.Context, userID string, enabled bool) error

	// 角色相关推送
	PushRoleCreated(ctx context.Context, roleID string, roleInfo interface{}) error
	PushRoleUpdated(ctx context.Context, roleID string, roleInfo interface{}) error
	PushRoleDeleted(ctx context.Context, roleID string) error
	PushRoleStatusChanged(ctx context.Context, roleID string, enabled bool) error

	// 系统通知推送
	PushSystemNotice(ctx context.Context, notice interface{}) error
	PushSystemAlert(ctx context.Context, alert interface{}) error
	PushSystemMaintenance(ctx context.Context, maintenance interface{}) error

	// 自定义事件推送
	PushCustomEvent(ctx context.Context, event libWebsocket.WebSocketEvent, data interface{}, targetType string, targetIDs []string) error
}

type messagePush struct{}

func RegisterMessagePushService() {
	messagePushInstance = &messagePush{}
}

// PushToAll 推送给所有客户端
func (s *messagePush) PushToAll(ctx context.Context, event libWebsocket.WebSocketEvent, data interface{}) error {
	response := &libWebsocket.WResponse{
		Event:     event,
		Data:      data,
		Code:      200,
		Timestamp: gtime.Now().Unix(),
	}
	libWebsocket.SendToAll(response)
	g.Log().Infof(ctx, "推送消息到所有客户端: event=%s", event)
	return nil
}

// PushToUser 推送给指定用户
func (s *messagePush) PushToUser(ctx context.Context, userID string, event libWebsocket.WebSocketEvent, data interface{}) error {
	response := &libWebsocket.WResponse{
		Event:     event,
		Data:      data,
		Code:      200,
		Timestamp: gtime.Now().Unix(),
	}
	libWebsocket.SendToUser(userID, response)
	g.Log().Infof(ctx, "推送消息到用户: userID=%s, event=%s", userID, event)
	return nil
}

// PushToUsers 推送给多个用户
func (s *messagePush) PushToUsers(ctx context.Context, userIDs []string, event libWebsocket.WebSocketEvent, data interface{}) error {
	for _, userID := range userIDs {
		if err := s.PushToUser(ctx, userID, event, data); err != nil {
			g.Log().Warningf(ctx, "推送消息到用户失败: userID=%s, event=%s, error=%v", userID, event, err)
		}
	}
	return nil
}

// PushToClient 推送给指定客户端
func (s *messagePush) PushToClient(ctx context.Context, clientID string, event libWebsocket.WebSocketEvent, data interface{}) error {
	response := &libWebsocket.WResponse{
		Event:     event,
		Data:      data,
		Code:      200,
		Timestamp: gtime.Now().Unix(),
	}
	libWebsocket.SendToClientID(clientID, response)
	g.Log().Infof(ctx, "推送消息到客户端: clientID=%s, event=%s", clientID, event)
	return nil
}

// PushToOrg 推送给指定组织
func (s *messagePush) PushToOrg(ctx context.Context, orgID string, event libWebsocket.WebSocketEvent, data interface{}) error {
	response := &libWebsocket.WResponse{
		Event:     event,
		Data:      data,
		Code:      200,
		Timestamp: gtime.Now().Unix(),
	}
	libWebsocket.SendToOrg(orgID, response)
	g.Log().Infof(ctx, "推送消息到组织: orgID=%s, event=%s", orgID, event)
	return nil
}

// PushOrgCreated 推送组织创建事件
func (s *messagePush) PushOrgCreated(ctx context.Context, orgID string, orgInfo interface{}) error {
	data := map[string]interface{}{
		"org_id":   orgID,
		"org_info": orgInfo,
		"action":   "created",
	}
	return s.PushToAll(ctx, "org_event", data)
}

// PushOrgUpdated 推送组织更新事件
func (s *messagePush) PushOrgUpdated(ctx context.Context, orgID string, orgInfo interface{}) error {
	data := map[string]interface{}{
		"org_id":   orgID,
		"org_info": orgInfo,
		"action":   "updated",
	}
	return s.PushToAll(ctx, "org_event", data)
}

// PushOrgDeleted 推送组织删除事件
func (s *messagePush) PushOrgDeleted(ctx context.Context, orgID string) error {
	data := map[string]interface{}{
		"org_id": orgID,
		"action": "deleted",
	}
	return s.PushToAll(ctx, "org_event", data)
}

// PushOrgStatusChanged 推送组织状态变更事件
func (s *messagePush) PushOrgStatusChanged(ctx context.Context, orgID string, enabled bool) error {
	data := map[string]interface{}{
		"org_id":  orgID,
		"enabled": enabled,
		"action":  "status_changed",
	}
	return s.PushToAll(ctx, "org_event", data)
}

// PushUserCreated 推送用户创建事件
func (s *messagePush) PushUserCreated(ctx context.Context, userID string, userInfo interface{}) error {
	data := map[string]interface{}{
		"user_id":   userID,
		"user_info": userInfo,
		"action":    "created",
	}
	return s.PushToAll(ctx, "user_event", data)
}

// PushUserUpdated 推送用户更新事件
func (s *messagePush) PushUserUpdated(ctx context.Context, userID string, userInfo interface{}) error {
	data := map[string]interface{}{
		"user_id":   userID,
		"user_info": userInfo,
		"action":    "updated",
	}
	return s.PushToAll(ctx, "user_event", data)
}

// PushUserDeleted 推送用户删除事件
func (s *messagePush) PushUserDeleted(ctx context.Context, userID string) error {
	data := map[string]interface{}{
		"user_id": userID,
		"action":  "deleted",
	}
	return s.PushToAll(ctx, "user_event", data)
}

// PushUserStatusChanged 推送用户状态变更事件
func (s *messagePush) PushUserStatusChanged(ctx context.Context, userID string, enabled bool) error {
	data := map[string]interface{}{
		"user_id": userID,
		"enabled": enabled,
		"action":  "status_changed",
	}
	return s.PushToAll(ctx, "user_event", data)
}

// PushRoleCreated 推送角色创建事件
func (s *messagePush) PushRoleCreated(ctx context.Context, roleID string, roleInfo interface{}) error {
	data := map[string]interface{}{
		"role_id":   roleID,
		"role_info": roleInfo,
		"action":    "created",
	}
	return s.PushToAll(ctx, "role_event", data)
}

// PushRoleUpdated 推送角色更新事件
func (s *messagePush) PushRoleUpdated(ctx context.Context, roleID string, roleInfo interface{}) error {
	data := map[string]interface{}{
		"role_id":   roleID,
		"role_info": roleInfo,
		"action":    "updated",
	}
	return s.PushToAll(ctx, "role_event", data)
}

// PushRoleDeleted 推送角色删除事件
func (s *messagePush) PushRoleDeleted(ctx context.Context, roleID string) error {
	data := map[string]interface{}{
		"role_id": roleID,
		"action":  "deleted",
	}
	return s.PushToAll(ctx, "role_event", data)
}

// PushRoleStatusChanged 推送角色状态变更事件
func (s *messagePush) PushRoleStatusChanged(ctx context.Context, roleID string, enabled bool) error {
	data := map[string]interface{}{
		"role_id": roleID,
		"enabled": enabled,
		"action":  "status_changed",
	}
	return s.PushToAll(ctx, "role_event", data)
}

// PushSystemNotice 推送系统通知
func (s *messagePush) PushSystemNotice(ctx context.Context, notice interface{}) error {
	data := map[string]interface{}{
		"type":   "notice",
		"notice": notice,
	}
	return s.PushToAll(ctx, "system_message", data)
}

// PushSystemAlert 推送系统告警
func (s *messagePush) PushSystemAlert(ctx context.Context, alert interface{}) error {
	data := map[string]interface{}{
		"type":  "alert",
		"alert": alert,
	}
	return s.PushToAll(ctx, "system_message", data)
}

// PushSystemMaintenance 推送系统维护通知
func (s *messagePush) PushSystemMaintenance(ctx context.Context, maintenance interface{}) error {
	data := map[string]interface{}{
		"type":        "maintenance",
		"maintenance": maintenance,
	}
	return s.PushToAll(ctx, "system_message", data)
}

// PushCustomEvent 推送自定义事件
func (s *messagePush) PushCustomEvent(ctx context.Context, event libWebsocket.WebSocketEvent, data interface{}, targetType string, targetIDs []string) error {
	switch targetType {
	case "all":
		return s.PushToAll(ctx, event, data)
	case "user":
		return s.PushToUsers(ctx, targetIDs, event, data)
	case "client":
		for _, clientID := range targetIDs {
			if err := s.PushToClient(ctx, clientID, event, data); err != nil {
				g.Log().Warningf(ctx, "推送自定义事件到客户端失败: clientID=%s, event=%s, error=%v", clientID, event, err)
			}
		}
		return nil
	default:
		return fmt.Errorf("不支持的目标类型: %s", targetType)
	}
}
