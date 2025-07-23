# WebSocket消息推送系统使用指南

## 概述

本系统基于GoFrame框架构建，提供了完整的WebSocket消息推送功能，支持实时向客户端推送各种类型的消息。

## 系统架构

### 核心组件

1. **WebSocket客户端管理器** (`library/libWebsocket`)
   - 客户端连接管理
   - 消息路由分发
   - 用户会话管理

2. **消息推送服务** (`internal/app/system/service/message_push.go`)
   - 统一的消息推送接口
   - 支持多种推送目标
   - 业务事件推送

3. **消息推送控制器** (`internal/app/system/controller/message_push.go`)
   - RESTful API接口
   - 消息推送管理
   - 客户端状态查询

4. **WebSocket消息处理器** (`internal/app/system/controller/websocket_handler.go`)
   - 客户端消息处理
   - 用户登录/登出
   - 标签管理

## 功能特性

### 1. 多种推送目标

- **全部客户端**: 向所有连接的客户端推送消息
- **指定用户**: 向特定用户的所有连接推送消息
- **指定客户端**: 向特定客户端推送消息
- **标签推送**: 向具有特定标签的客户端推送消息

### 2. 业务事件推送

- **组织事件**: 组织创建、更新、删除、状态变更
- **用户事件**: 用户创建、更新、删除、状态变更
- **角色事件**: 角色创建、更新、删除、状态变更

### 3. 系统消息推送

- **系统通知**: 一般性系统通知
- **系统告警**: 系统告警信息
- **维护通知**: 系统维护相关信息

### 4. 客户端管理

- **在线状态查询**: 查询用户在线状态
- **客户端信息**: 获取客户端详细信息
- **标签管理**: 添加/移除客户端标签

## API接口

### 基础推送接口

#### 1. 推送给所有客户端
```http
POST /message/push/all
Content-Type: application/json

{
    "event": "system_notice",
    "data": {
        "title": "系统通知",
        "content": "这是一条系统通知"
    }
}
```

#### 2. 推送给指定用户
```http
POST /message/push/user
Content-Type: application/json

{
    "user_id": "user_001",
    "event": "user_message",
    "data": {
        "message": "您有一条新消息"
    }
}
```

#### 3. 推送给多个用户
```http
POST /message/push/users
Content-Type: application/json

{
    "user_ids": ["user_001", "user_002", "user_003"],
    "event": "batch_notice",
    "data": {
        "title": "批量通知",
        "content": "这是一条批量通知"
    }
}
```

#### 4. 推送给指定客户端
```http
POST /message/push/client
Content-Type: application/json

{
    "client_id": "client_001",
    "event": "direct_message",
    "data": {
        "message": "直接消息"
    }
}
```

#### 5. 推送给指定标签
```http
POST /message/push/tag
Content-Type: application/json

{
    "tag": "admin",
    "event": "admin_notice",
    "data": {
        "title": "管理员通知",
        "content": "这是一条管理员通知"
    }
}
```

### 业务事件推送接口

#### 1. 推送组织事件
```http
POST /message/push/org
Content-Type: application/json

{
    "org_id": "org_001",
    "action": "created",
    "org_info": {
        "id": "org_001",
        "name": "技术部",
        "code": "TECH"
    }
}
```

支持的动作：
- `created`: 组织创建
- `updated`: 组织更新
- `deleted`: 组织删除
- `status_changed`: 组织状态变更

#### 2. 推送用户事件
```http
POST /message/push/user-event
Content-Type: application/json

{
    "user_id": "user_001",
    "action": "updated",
    "user_info": {
        "id": "user_001",
        "name": "张三",
        "email": "zhangsan@example.com"
    }
}
```

#### 3. 推送角色事件
```http
POST /message/push/role
Content-Type: application/json

{
    "role_id": "role_001",
    "action": "created",
    "role_info": {
        "id": "role_001",
        "name": "管理员",
        "code": "ADMIN"
    }
}
```

### 系统消息推送接口

#### 1. 推送系统通知
```http
POST /message/push/system/notice
Content-Type: application/json

{
    "title": "系统维护通知",
    "content": "系统将于今晚22:00进行维护",
    "level": "info",
    "data": {
        "maintenance_time": "2024-01-15 22:00:00"
    }
}
```

#### 2. 推送系统告警
```http
POST /message/push/system/alert
Content-Type: application/json

{
    "title": "系统告警",
    "content": "服务器CPU使用率超过90%",
    "level": "warning",
    "data": {
        "cpu_usage": 95,
        "server_id": "server_001"
    }
}
```

#### 3. 推送系统维护通知
```http
POST /message/push/system/maintenance
Content-Type: application/json

{
    "title": "系统维护",
    "content": "系统将进行例行维护",
    "start_time": "2024-01-15 22:00:00",
    "end_time": "2024-01-16 02:00:00",
    "affected_services": ["用户管理", "组织管理"],
    "data": {
        "maintenance_type": "routine"
    }
}
```

### 自定义事件推送接口

```http
POST /message/push/custom
Content-Type: application/json

{
    "event": "custom_event",
    "data": {
        "custom_field": "custom_value"
    },
    "target_type": "user",
    "target_ids": ["user_001", "user_002"]
}
```

支持的目标类型：
- `all`: 所有客户端
- `user`: 指定用户
- `client`: 指定客户端
- `tag`: 指定标签

### 客户端管理接口

#### 1. 获取在线客户端列表
```http
GET /message/clients?page=1&page_size=20&user_id=user_001&tag=admin
```

#### 2. 获取客户端统计信息
```http
GET /message/clients/stats
```

## WebSocket客户端连接

### 连接地址
```
ws://your-domain/websocket/
```

### 客户端消息格式

#### 1. 用户登录
```json
{
    "event": "login",
    "data": {
        "user_id": "user_001",
        "user_name": "张三"
    }
}
```

#### 2. 心跳
```json
{
    "event": "heartbeat",
    "data": {}
}
```

#### 3. 订阅事件
```json
{
    "event": "subscribe",
    "data": {
        "event": "org_event"
    }
}
```

#### 4. 取消订阅
```json
{
    "event": "unsubscribe",
    "data": {
        "event": "org_event"
    }
}
```

#### 5. 添加标签
```json
{
    "event": "add_tag",
    "data": {
        "tag": "admin"
    }
}
```

#### 6. 移除标签
```json
{
    "event": "remove_tag",
    "data": {
        "tag": "admin"
    }
}
```

#### 7. 获取客户端信息
```json
{
    "event": "get_client_info",
    "data": {}
}
```

#### 8. 获取在线状态
```json
{
    "event": "get_online_status",
    "data": {
        "user_id": "user_001"
    }
}
```

### 服务器响应格式

#### 成功响应
```json
{
    "event": "login",
    "data": {
        "user_id": "user_001",
        "login_time": 1640995200,
        "client_id": "client_001"
    },
    "code": 200,
    "timestamp": 1640995200
}
```

#### 错误响应
```json
{
    "event": "login",
    "code": 0,
    "errorMsg": "用户ID不能为空",
    "timestamp": 1640995200
}
```

## 使用示例

### 1. 在业务代码中推送消息

```go
// 推送组织创建事件
err := service.MessagePush().PushOrgCreated(ctx, "org_001", orgInfo)
if err != nil {
    g.Log().Error(ctx, "推送组织创建事件失败:", err)
}

// 推送系统通知
notice := map[string]interface{}{
    "title":   "系统通知",
    "content": "这是一条系统通知",
    "level":   "info",
}
err = service.MessagePush().PushSystemNotice(ctx, notice)
if err != nil {
    g.Log().Error(ctx, "推送系统通知失败:", err)
}

// 推送给指定用户
err = service.MessagePush().PushToUser(ctx, "user_001", "user_message", map[string]interface{}{
    "message": "您有一条新消息",
})
if err != nil {
    g.Log().Error(ctx, "推送用户消息失败:", err)
}
```

### 2. 在组织管理中使用推送

```go
// 组织创建后推送事件
func (c *orgController) Add(ctx context.Context, req *system.OrgAddReq) (res *system.OrgAddRes, err error) {
    // ... 创建组织的逻辑 ...
    
    // 推送组织创建事件
    err = service.MessagePush().PushOrgCreated(ctx, orgID, orgInfo)
    if err != nil {
        g.Log().Warning(ctx, "推送组织创建事件失败:", err)
    }
    
    return
}
```

### 3. JavaScript客户端示例

```javascript
// 连接WebSocket
const ws = new WebSocket('ws://your-domain/websocket/');

// 用户登录
ws.send(JSON.stringify({
    event: 'login',
    data: {
        user_id: 'user_001',
        user_name: '张三'
    }
}));

// 订阅组织事件
ws.send(JSON.stringify({
    event: 'subscribe',
    data: {
        event: 'org_event'
    }
}));

// 添加标签
ws.send(JSON.stringify({
    event: 'add_tag',
    data: {
        tag: 'admin'
    }
}));

// 监听消息
ws.onmessage = function(event) {
    const message = JSON.parse(event.data);
    
    switch(message.event) {
        case 'org_event':
            handleOrgEvent(message.data);
            break;
        case 'system_message':
            handleSystemMessage(message.data);
            break;
        case 'user_message':
            handleUserMessage(message.data);
            break;
        default:
            console.log('收到消息:', message);
    }
};

// 处理组织事件
function handleOrgEvent(data) {
    switch(data.action) {
        case 'created':
            console.log('组织创建:', data.org_info);
            break;
        case 'updated':
            console.log('组织更新:', data.org_info);
            break;
        case 'deleted':
            console.log('组织删除:', data.org_id);
            break;
        case 'status_changed':
            console.log('组织状态变更:', data.org_id, data.enabled);
            break;
    }
}

// 处理系统消息
function handleSystemMessage(data) {
    switch(data.type) {
        case 'notice':
            showNotification(data.notice.title, data.notice.content, 'info');
            break;
        case 'alert':
            showNotification(data.alert.title, data.alert.content, 'warning');
            break;
        case 'maintenance':
            showMaintenanceNotice(data.maintenance);
            break;
    }
}
```

## 配置说明

### WebSocket配置

在 `config.yaml` 中添加WebSocket配置：

```yaml
websocket:
  # WebSocket端口
  port: 8080
  # 心跳超时时间（秒）
  heartbeat_timeout: 300
  # 最大连接数
  max_connections: 10000
  # 消息缓冲区大小
  message_buffer_size: 1000
```

### 消息推送配置

```yaml
message_push:
  # 是否启用消息推送
  enabled: true
  # 推送重试次数
  retry_count: 3
  # 推送超时时间（秒）
  timeout: 30
  # 批量推送大小
  batch_size: 100
```

## 注意事项

1. **连接管理**: 客户端需要定期发送心跳消息保持连接
2. **错误处理**: 客户端需要处理连接断开和重连
3. **消息去重**: 避免重复推送相同消息
4. **性能优化**: 大量消息推送时考虑使用批量推送
5. **安全考虑**: 验证用户权限，防止未授权推送
6. **日志记录**: 记录重要的推送事件和错误信息

## 故障排除

### 常见问题

1. **连接失败**
   - 检查WebSocket服务是否启动
   - 检查防火墙设置
   - 检查网络连接

2. **消息推送失败**
   - 检查目标用户是否在线
   - 检查消息格式是否正确
   - 查看服务器日志

3. **客户端断开**
   - 检查心跳超时设置
   - 检查网络稳定性
   - 实现自动重连机制

### 调试工具

1. **WebSocket调试**: 使用浏览器开发者工具
2. **服务器日志**: 查看详细的推送日志
3. **客户端状态**: 使用API查询客户端状态

## 扩展功能

1. **消息持久化**: 将消息存储到数据库
2. **消息历史**: 提供消息历史查询功能
3. **消息统计**: 统计推送消息的数量和成功率
4. **消息模板**: 支持消息模板功能
5. **消息优先级**: 支持消息优先级设置
6. **消息过滤**: 支持消息内容过滤 