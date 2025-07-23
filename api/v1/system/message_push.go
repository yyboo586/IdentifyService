package system

import (
	"IdentifyService/internal/app/system/model"
	"IdentifyService/library/libWebsocket"

	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 基础推送接口 ====================

// MessagePushToAllReq 推送给所有客户端
type MessagePushToAllReq struct {
	g.Meta `path:"/message/push/all" tags:"消息推送" method:"post" summary:"推送给所有客户端"`
	model.Author
	Event libWebsocket.WebSocketEvent `json:"event" v:"required#事件名称不能为空" dc:"事件名称"`
	Data  interface{}                 `json:"data" dc:"推送数据"`
}

type MessagePushToAllRes struct {
	g.Meta `mime:"application/json"`
}

// MessagePushToUserReq 推送给指定用户
type MessagePushToUserReq struct {
	g.Meta `path:"/message/push/user" tags:"消息推送" method:"post" summary:"推送给指定用户"`
	model.Author
	UserID string                      `json:"user_id" v:"required#用户ID不能为空" dc:"用户ID"`
	Event  libWebsocket.WebSocketEvent `json:"event" v:"required#事件名称不能为空" dc:"事件名称"`
	Data   interface{}                 `json:"data" dc:"推送数据"`
}

type MessagePushToUserRes struct {
	g.Meta `mime:"application/json"`
}

// MessagePushToUsersReq 推送给多个用户
type MessagePushToUsersReq struct {
	g.Meta `path:"/message/push/users" tags:"消息推送" method:"post" summary:"推送给多个用户"`
	model.Author
	UserIDs []string                    `json:"user_ids" v:"required#用户ID列表不能为空" dc:"用户ID列表"`
	Event   libWebsocket.WebSocketEvent `json:"event" v:"required#事件名称不能为空" dc:"事件名称"`
	Data    interface{}                 `json:"data" dc:"推送数据"`
}

type MessagePushToUsersRes struct {
	g.Meta `mime:"application/json"`
}

// MessagePushToClientReq 推送给指定客户端
type MessagePushToClientReq struct {
	g.Meta `path:"/message/push/client" tags:"消息推送" method:"post" summary:"推送给指定客户端"`
	model.Author
	ClientID string                      `json:"client_id" v:"required#客户端ID不能为空" dc:"客户端ID"`
	Event    libWebsocket.WebSocketEvent `json:"event" v:"required#事件名称不能为空" dc:"事件名称"`
	Data     interface{}                 `json:"data" dc:"推送数据"`
}

type MessagePushToClientRes struct {
	g.Meta `mime:"application/json"`
}

// MessagePushToTagReq 推送给指定标签
type MessagePushToTagReq struct {
	g.Meta `path:"/message/push/tag" tags:"消息推送" method:"post" summary:"推送给指定标签"`
	model.Author
	Tag   string                      `json:"tag" v:"required#标签不能为空" dc:"标签"`
	Event libWebsocket.WebSocketEvent `json:"event" v:"required#事件名称不能为空" dc:"事件名称"`
	Data  interface{}                 `json:"data" dc:"推送数据"`
}

type MessagePushToTagRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 业务事件推送接口 ====================

// MessagePushOrgEventReq 推送组织事件
type MessagePushOrgEventReq struct {
	g.Meta `path:"/message/push/org" tags:"消息推送" method:"post" summary:"推送组织事件"`
	model.Author
	OrgID   string      `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Action  string      `json:"action" v:"required#动作不能为空" dc:"动作(created/updated/deleted/status_changed)"`
	OrgInfo interface{} `json:"org_info" dc:"组织信息"`
	Enabled *bool       `json:"enabled" dc:"状态(仅status_changed时使用)"`
}

type MessagePushOrgEventRes struct {
	g.Meta `mime:"application/json"`
}

// MessagePushUserEventReq 推送用户事件
type MessagePushUserEventReq struct {
	g.Meta `path:"/message/push/user-event" tags:"消息推送" method:"post" summary:"推送用户事件"`
	model.Author
	UserID   string      `json:"user_id" v:"required#用户ID不能为空" dc:"用户ID"`
	Action   string      `json:"action" v:"required#动作不能为空" dc:"动作(created/updated/deleted/status_changed)"`
	UserInfo interface{} `json:"user_info" dc:"用户信息"`
	Enabled  *bool       `json:"enabled" dc:"状态(仅status_changed时使用)"`
}

type MessagePushUserEventRes struct {
	g.Meta `mime:"application/json"`
}

// MessagePushRoleEventReq 推送角色事件
type MessagePushRoleEventReq struct {
	g.Meta `path:"/message/push/role" tags:"消息推送" method:"post" summary:"推送角色事件"`
	model.Author
	RoleID   string      `json:"role_id" v:"required#角色ID不能为空" dc:"角色ID"`
	Action   string      `json:"action" v:"required#动作不能为空" dc:"动作(created/updated/deleted/status_changed)"`
	RoleInfo interface{} `json:"role_info" dc:"角色信息"`
	Enabled  *bool       `json:"enabled" dc:"状态(仅status_changed时使用)"`
}

type MessagePushRoleEventRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 系统消息推送接口 ====================

// MessagePushSystemNoticeReq 推送系统通知
type MessagePushSystemNoticeReq struct {
	g.Meta `path:"/message/push/system/notice" tags:"消息推送" method:"post" summary:"推送系统通知"`
	model.Author
	Title   string      `json:"title" v:"required#标题不能为空" dc:"通知标题"`
	Content string      `json:"content" v:"required#内容不能为空" dc:"通知内容"`
	Level   string      `json:"level" d:"info" dc:"通知级别(info/warning/error)"`
	Data    interface{} `json:"data" dc:"额外数据"`
}

type MessagePushSystemNoticeRes struct {
	g.Meta `mime:"application/json"`
}

// MessagePushSystemAlertReq 推送系统告警
type MessagePushSystemAlertReq struct {
	g.Meta `path:"/message/push/system/alert" tags:"消息推送" method:"post" summary:"推送系统告警"`
	model.Author
	Title   string      `json:"title" v:"required#标题不能为空" dc:"告警标题"`
	Content string      `json:"content" v:"required#内容不能为空" dc:"告警内容"`
	Level   string      `json:"level" d:"warning" dc:"告警级别(warning/error/critical)"`
	Data    interface{} `json:"data" dc:"额外数据"`
}

type MessagePushSystemAlertRes struct {
	g.Meta `mime:"application/json"`
}

// MessagePushSystemMaintenanceReq 推送系统维护通知
type MessagePushSystemMaintenanceReq struct {
	g.Meta `path:"/message/push/system/maintenance" tags:"消息推送" method:"post" summary:"推送系统维护通知"`
	model.Author
	Title            string      `json:"title" v:"required#标题不能为空" dc:"维护标题"`
	Content          string      `json:"content" v:"required#内容不能为空" dc:"维护内容"`
	StartTime        string      `json:"start_time" v:"required#开始时间不能为空" dc:"开始时间"`
	EndTime          string      `json:"end_time" v:"required#结束时间不能为空" dc:"结束时间"`
	AffectedServices []string    `json:"affected_services" dc:"受影响的服务"`
	Data             interface{} `json:"data" dc:"额外数据"`
}

type MessagePushSystemMaintenanceRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 自定义事件推送接口 ====================

// MessagePushCustomEventReq 推送自定义事件
type MessagePushCustomEventReq struct {
	g.Meta `path:"/message/push/custom" tags:"消息推送" method:"post" summary:"推送自定义事件"`
	model.Author
	Event      libWebsocket.WebSocketEvent `json:"event" v:"required#事件名称不能为空" dc:"事件名称"`
	Data       interface{}                 `json:"data" dc:"事件数据"`
	TargetType string                      `json:"target_type" v:"required#目标类型不能为空" dc:"目标类型(all/user/client/tag)"`
	TargetIDs  []string                    `json:"target_ids" dc:"目标ID列表(非all时必填)"`
}

type MessagePushCustomEventRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 客户端管理接口 ====================

// MessageGetClientsReq 获取在线客户端列表
type MessageGetClientsReq struct {
	g.Meta `path:"/message/clients" tags:"消息推送" method:"get" summary:"获取在线客户端列表"`
	model.Author
	Page     int    `p:"page" d:"1" dc:"页码"`
	PageSize int    `p:"page_size" d:"20" dc:"每页数量"`
	UserID   string `p:"user_id" dc:"用户ID过滤"`
	Tag      string `p:"tag" dc:"标签过滤"`
}

type MessageGetClientsRes struct {
	g.Meta   `mime:"application/json"`
	List     []*ClientInfo `json:"list" dc:"客户端列表"`
	Total    int           `json:"total" dc:"总数"`
	Page     int           `json:"page" dc:"当前页"`
	PageSize int           `json:"page_size" dc:"每页数量"`
}

// MessageGetClientStatsReq 获取客户端统计信息
type MessageGetClientStatsReq struct {
	g.Meta `path:"/message/clients/stats" tags:"消息推送" method:"get" summary:"获取客户端统计信息"`
	model.Author
}

type MessageGetClientStatsRes struct {
	g.Meta `mime:"application/json"`
	*ClientStatsInfo
}

// ==================== 数据模型 ====================

// ClientInfo 客户端信息
type ClientInfo struct {
	ID            string `json:"id" dc:"客户端ID"`
	UserID        string `json:"user_id" dc:"用户ID"`
	UserName      string `json:"user_name" dc:"用户名称"`
	IP            string `json:"ip" dc:"客户端IP"`
	UserAgent     string `json:"user_agent" dc:"用户代理"`
	ConnectTime   int64  `json:"connect_time" dc:"连接时间"`
	HeartbeatTime int64  `json:"heartbeat_time" dc:"心跳时间"`
}

// ClientStatsInfo 客户端统计信息
type ClientStatsInfo struct {
	TotalClients    int `json:"total_clients" dc:"总客户端数"`
	OnlineUsers     int `json:"online_users" dc:"在线用户数"`
	TotalUsers      int `json:"total_users" dc:"总用户数"`
	ActiveClients   int `json:"active_clients" dc:"活跃客户端数"`
	InactiveClients int `json:"inactive_clients" dc:"非活跃客户端数"`
}
