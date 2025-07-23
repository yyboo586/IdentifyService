# Casbin权限管理框架详解

## 1. Casbin核心概念

### 1.1 什么是Casbin
Casbin是一个强大的权限管理库，使用策略语言来定义权限规则。它支持多种权限模型，包括RBAC（基于角色的访问控制）。

### 1.2 RBAC模型配置 (`rbac_model.conf`)

```conf
[request_definition]
r = sub, obj, act    # 请求定义：主体(用户/角色)、客体(资源)、操作

[policy_definition]  
p = sub, obj, act    # 策略定义：角色、资源、操作

[role_definition]
g = _, _             # 角色定义：用户、角色

[policy_effect]
e = some(where (p.eft == allow))  # 策略效果：允许访问

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act  # 匹配器：检查权限
```

**配置说明：**
- `r = sub, obj, act`: 定义请求格式，包含主体、客体、操作
- `p = sub, obj, act`: 定义策略格式，包含角色、资源、操作
- `g = _, _`: 定义角色继承关系，用户继承角色
- `e = some(where (p.eft == allow))`: 定义策略效果，只要有一个策略允许就通过
- `m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act`: 定义匹配规则

## 2. Enforcer核心方法详解

### 2.1 权限检查方法

```go
// Enforce 检查权限
// 参数：角色ID, 资源ID, 操作类型
// 返回：true表示有权限，false表示无权限
enforcer.Enforce("角色ID", "资源ID", "操作类型")

// 示例：
enforcer.Enforce("1", "100", "All")  // 检查角色1是否有对资源100的所有操作权限
```

### 2.2 策略管理方法

#### **添加策略**
```go
// AddPolicies 批量添加权限策略
// 参数：策略列表，每个策略格式为 [角色ID, 资源ID, 操作]
rules := [][]string{
    {"1", "100", "All"},  // 角色1对资源100有所有权限
    {"2", "200", "Read"}, // 角色2对资源200有读权限
}
enforcer.AddPolicies(rules)

// AddNamedPolicies 添加命名策略
// 参数：策略类型("p"表示权限策略), 策略列表
enforcer.AddNamedPolicies("p", rules)
```

#### **删除策略**
```go
// RemoveFilteredPolicy 删除过滤的策略
// 参数：字段索引, 字段值
enforcer.RemoveFilteredPolicy(0, "角色ID")  // 删除指定角色的所有权限

// RemovePolicy 删除指定策略
// 参数：策略 [角色ID, 资源ID, 操作]
enforcer.RemovePolicy([]string{"1", "100", "All"})
```

#### **查询策略**
```go
// GetFilteredPolicy 获取过滤的策略
// 参数：字段索引, 字段值
policies := enforcer.GetFilteredPolicy(0, "角色ID")  // 获取指定角色的所有权限

// GetFilteredNamedPolicy 获取过滤的命名策略
// 参数：策略类型, 字段索引, 字段值
policies := enforcer.GetFilteredNamedPolicy("p", 0, "角色ID")
```

### 2.3 角色管理方法

```go
// GetFilteredGroupingPolicy 获取过滤的分组策略（用户-角色关系）
// 参数：字段索引, 字段值
userRoles := enforcer.GetFilteredGroupingPolicy(1, "角色ID")  // 获取指定角色的所有用户

// AddGroupingPolicy 添加用户-角色关系
enforcer.AddGroupingPolicy("u_100", "1")  // 用户100分配角色1

// RemoveGroupingPolicy 删除用户-角色关系
enforcer.RemoveGroupingPolicy("u_100", "1")  // 删除用户100的角色1
```

## 3. 代码中的应用示例

### 3.1 为角色分配权限

```go
func addRoleRule(ctx context.Context, roleID int64, ruleIDs []int64) error {
    enforcer, _ := service.CasbinEnforcer(ctx)
    
    // 构建策略列表
    rules := make([][]string, len(ruleIDs))
    for k, v := range ruleIDs {
        // 格式：[角色ID, 资源ID, 操作]
        rules[k] = []string{gconv.String(roleID), gconv.String(v), "All"}
    }
    
    // 批量添加权限策略
    _, err := enforcer.AddNamedPolicies("p", rules)
    return err
}
```

### 3.2 删除角色权限

```go
func delRoleRule(ctx context.Context, roleIDs []int64) error {
    enforcer, _ := service.CasbinEnforcer(ctx)
    
    for _, roleID := range roleIDs {
        // 删除指定角色的所有权限策略
        // 参数：0表示第一个字段（角色ID），roleID是要删除的角色ID
        _, err := enforcer.RemoveFilteredPolicy(0, gconv.String(roleID))
        if err != nil {
            return err
        }
    }
    return nil
}
```

### 3.3 获取角色权限

```go
func getRolePermissions(ctx context.Context, roleID int64) ([]int64, error) {
    enforcer, _ := service.CasbinEnforcer(ctx)
    
    // 获取角色的所有权限策略
    policies := enforcer.GetFilteredNamedPolicy("p", 0, gconv.String(roleID))
    
    // 提取资源ID（第二个字段）
    var resourceIDs []int64
    for _, policy := range policies {
        resourceID := gconv.Int64(policy[1])  // policy[1]是资源ID
        resourceIDs = append(resourceIDs, resourceID)
    }
    
    return resourceIDs, nil
}
```

### 3.4 权限检查

```go
func checkPermission(ctx context.Context, roleID int64, resourceID int64) bool {
    enforcer, _ := service.CasbinEnforcer(ctx)
    
    // 检查角色是否有对资源的操作权限
    hasPermission, _ := enforcer.Enforce(
        gconv.String(roleID),    // 角色ID
        gconv.String(resourceID), // 资源ID
        "All"                     // 操作类型
    )
    
    return hasPermission
}
```

## 4. 数据存储格式

### 4.1 权限策略表 (t_casbin_rule)

| 字段 | 说明 | 示例 |
|------|------|------|
| ptype | 策略类型 | "p" (权限策略) |
| v0 | 第一个值 | "1" (角色ID) |
| v1 | 第二个值 | "100" (资源ID) |
| v2 | 第三个值 | "All" (操作类型) |

**示例数据：**
```
p, 1, 100, All    # 角色1对资源100有所有权限
p, 2, 200, Read   # 角色2对资源200有读权限
```

### 4.2 用户角色关系表

| 字段 | 说明 | 示例 |
|------|------|------|
| ptype | 策略类型 | "g" (分组策略) |
| v0 | 用户ID | "u_100" (用户100) |
| v1 | 角色ID | "1" (角色1) |

**示例数据：**
```
g, u_100, 1    # 用户100分配角色1
g, u_200, 2    # 用户200分配角色2
```

## 5. 权限检查流程

### 5.1 完整权限检查流程

```go
func checkAuth(ctx context.Context, userID string, resourceID int64) bool {
    // 1. 获取用户的所有角色
    roleIDs := getUserRoleIDs(userID)
    
    // 2. 遍历每个角色，检查权限
    enforcer, _ := service.CasbinEnforcer(ctx)
    for _, roleID := range roleIDs {
        hasPermission, _ := enforcer.Enforce(
            gconv.String(roleID),
            gconv.String(resourceID),
            "All"
        )
        if hasPermission {
            return true  // 只要有一个角色有权限就通过
        }
    }
    
    return false  // 所有角色都没有权限
}
```

### 5.2 权限继承机制

Casbin支持角色继承，通过`g`策略定义：

```go
// 添加角色继承关系
enforcer.AddGroupingPolicy("admin", "user")     // admin继承user角色
enforcer.AddGroupingPolicy("super_admin", "admin")  // super_admin继承admin角色

// 检查权限时会自动检查继承的角色
hasPermission, _ := enforcer.Enforce("super_admin", "resource", "read")
// 如果super_admin没有直接权限，会检查admin和user的权限
```

## 6. 最佳实践

### 6.1 权限策略设计

1. **最小权限原则**：只给角色分配必要的权限
2. **角色层次化**：设计合理的角色继承关系
3. **权限粒度控制**：根据业务需求设置合适的权限粒度

### 6.2 性能优化

1. **缓存策略**：对频繁查询的权限结果进行缓存
2. **批量操作**：使用批量API减少数据库访问次数
3. **索引优化**：为权限表添加合适的索引

### 6.3 安全考虑

1. **权限验证**：在关键操作前进行权限检查
2. **审计日志**：记录权限变更和访问日志
3. **定期清理**：定期清理无效的权限策略

## 7. 常见问题

### 7.1 权限不生效
- 检查策略是否正确添加到数据库
- 确认Enforcer是否正确加载了策略
- 验证权限检查的参数格式

### 7.2 性能问题
- 使用缓存减少权限检查次数
- 优化数据库查询
- 考虑使用Redis等高性能存储

### 7.3 权限继承问题
- 确认角色继承关系是否正确设置
- 检查`g`策略是否正确配置
- 验证继承链是否完整 