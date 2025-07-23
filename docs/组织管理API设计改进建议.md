# 组织管理API设计改进建议

## 当前设计问题分析

### 1. API路径设计不够RESTful

**当前设计：**
```go
POST   /org                    // 创建组织
DELETE /org/{id}              // 删除组织  
PUT    /org/{id}              // 更新基本信息
PUT    /org/{id}/status       // 更新状态
GET    /org/{id}              // 获取详情
GET    /org/{id}/tree         // 获取树形结构
GET    /org/trees             // 获取所有树形结构
```

**问题：**
- 路径不一致：`/org` vs `/org/trees`
- 状态更新使用PUT而不是PATCH
- 缺少复数形式：应该使用`/orgs`而不是`/org`

**改进建议：**
```go
POST   /orgs                  // 创建组织
GET    /orgs                  // 获取组织列表
GET    /orgs/{id}            // 获取组织详情
PUT    /orgs/{id}            // 更新组织信息
PATCH  /orgs/{id}/status     // 更新组织状态
DELETE /orgs/{id}            // 删除组织
GET    /orgs/tree            // 获取组织树
GET    /orgs/{id}/subtree    // 获取子组织树
PATCH  /orgs/{id}/move       // 移动组织
POST   /orgs/{id}/copy       // 复制组织
GET    /orgs/{id}/stats      // 获取统计信息
DELETE /orgs/batch           // 批量删除
PATCH  /orgs/batch/status    // 批量更新状态
```

### 2. 请求/响应结构设计问题

**当前问题：**
- 缺少分页支持
- 缺少搜索和过滤功能
- 响应结构不够统一
- 缺少批量操作支持
- 缺少组织编码字段
- 缺少描述和排序字段

**改进建议：**

```go
// 组织列表请求
type OrgListReq struct {
    Page     int    `p:"page" d:"1" dc:"页码"`
    PageSize int    `p:"page_size" d:"20" dc:"每页数量"`
    Keyword  string `p:"keyword" dc:"搜索关键词"`
    PID      string `p:"pid" dc:"父级ID过滤"`
    Enabled  *bool  `p:"enabled" dc:"状态过滤"`
    SortBy   string `p:"sort_by" d:"created_at" dc:"排序字段"`
    SortDesc bool   `p:"sort_desc" d:"true" dc:"是否降序"`
}

// 组织列表响应
type OrgListRes struct {
    List     []*OrgInfo `json:"list" dc:"组织列表"`
    Total    int        `json:"total" dc:"总数"`
    Page     int        `json:"page" dc:"当前页"`
    PageSize int        `json:"page_size" dc:"每页数量"`
}
```

### 3. 业务逻辑覆盖不完整

**缺少的功能：**
- 组织移动功能
- 组织复制功能
- 组织统计信息
- 批量操作
- 组织编码管理
- 组织路径和层级信息

**新增功能建议：**

```go
// 组织移动
PATCH /orgs/{id}/move
{
    "new_pid": "新的父级ID"
}

// 组织复制
POST /orgs/{id}/copy
{
    "new_pid": "目标父级ID",
    "new_name": "新组织名称",
    "copy_users": true,
    "copy_roles": true,
    "copy_perms": true
}

// 组织统计
GET /orgs/{id}/stats
{
    "total_users": 100,
    "active_users": 80,
    "total_roles": 10,
    "total_children": 5,
    "direct_children": 3,
    "max_depth": 3
}
```

## 改进后的完整API设计

### 1. 基础CRUD操作

```go
// 创建组织
POST /orgs
{
    "name": "技术部",
    "pid": "parent_id",
    "code": "TECH",
    "manager_id": "user_id",
    "manager_name": "张三",
    "description": "负责技术开发",
    "sort": 1
}

// 获取组织列表
GET /orgs?page=1&page_size=20&keyword=技术&enabled=true

// 获取组织详情
GET /orgs/{id}

// 更新组织信息
PUT /orgs/{id}
{
    "name": "技术研发部",
    "code": "TECH_DEV",
    "manager_id": "new_user_id",
    "manager_name": "李四",
    "description": "负责技术研发工作",
    "sort": 2
}

// 删除组织
DELETE /orgs/{id}
```

### 2. 状态管理

```go
// 更新组织状态
PATCH /orgs/{id}/status
{
    "enabled": false
}

// 批量更新状态
PATCH /orgs/batch/status
{
    "ids": ["id1", "id2", "id3"],
    "enabled": true
}
```

### 3. 树形结构操作

```go
// 获取组织树
GET /orgs/tree?root_id=root&max_depth=5&enabled=true

// 获取子组织树
GET /orgs/{id}/subtree?max_depth=3&enabled=true
```

### 4. 高级操作

```go
// 移动组织
PATCH /orgs/{id}/move
{
    "new_pid": "new_parent_id"
}

// 复制组织
POST /orgs/{id}/copy
{
    "new_pid": "target_parent_id",
    "new_name": "技术部-副本",
    "copy_users": true,
    "copy_roles": true,
    "copy_perms": false
}

// 获取统计信息
GET /orgs/{id}/stats
```

### 5. 批量操作

```go
// 批量删除
DELETE /orgs/batch
{
    "ids": ["id1", "id2", "id3"]
}

// 批量更新状态
PATCH /orgs/batch/status
{
    "ids": ["id1", "id2", "id3"],
    "enabled": false
}
```

## 数据模型改进

### 1. 组织基本信息

```go
type OrgInfo struct {
    ID          string      `json:"id" dc:"组织ID"`
    PID         string      `json:"pid" dc:"父级ID"`
    Name        string      `json:"name" dc:"组织名称"`
    Code        string      `json:"code" dc:"组织编码"`
    ManagerID   string      `json:"manager_id" dc:"负责人ID"`
    ManagerName string      `json:"manager_name" dc:"负责人名称"`
    Description string      `json:"description" dc:"组织描述"`
    Enabled     bool        `json:"enabled" dc:"组织状态"`
    Sort        int         `json:"sort" dc:"排序"`
    CreatedAt   *gtime.Time `json:"created_at" dc:"创建时间"`
    UpdatedAt   *gtime.Time `json:"updated_at" dc:"修改时间"`
}
```

### 2. 组织详细信息

```go
type OrgDetailInfo struct {
    *OrgInfo
    UserCount     int    `json:"user_count" dc:"用户数量"`
    RoleCount     int    `json:"role_count" dc:"角色数量"`
    ChildrenCount int    `json:"children_count" dc:"子组织数量"`
    Level         int    `json:"level" dc:"组织层级"`
    Path          string `json:"path" dc:"组织路径"`
}
```

### 3. 组织树节点

```go
type OrgTreeNode struct {
    *OrgInfo
    Children []*OrgTreeNode `json:"children" dc:"子组织"`
    Level    int            `json:"level" dc:"层级"`
    Path     string         `json:"path" dc:"路径"`
}
```

## 实现建议

### 1. 数据库设计改进

```sql
ALTER TABLE t_org ADD COLUMN code VARCHAR(50) COMMENT '组织编码';
ALTER TABLE t_org ADD COLUMN description TEXT COMMENT '组织描述';
ALTER TABLE t_org ADD COLUMN sort INT DEFAULT 0 COMMENT '排序';
ALTER TABLE t_org ADD COLUMN level INT DEFAULT 1 COMMENT '组织层级';
ALTER TABLE t_org ADD COLUMN path VARCHAR(500) COMMENT '组织路径';
ALTER TABLE t_org ADD UNIQUE KEY uk_code (code);
```

### 2. 服务层接口扩展

```go
type IOrg interface {
    // 基础操作
    Create(ctx context.Context, in *model.Org) (orgID string, err error)
    GetDetail(ctx context.Context, id string) (*system.OrgDetailInfo, error)
    Update(ctx context.Context, in *model.Org) error
    Delete(ctx context.Context, id string) error
    UpdateStatus(ctx context.Context, id string, enabled bool) error
    
    // 列表和搜索
    List(ctx context.Context, params *model.OrgListParams) (list []*system.OrgInfo, total int, err error)
    
    // 树形结构
    GetTree(ctx context.Context, params *model.OrgTreeParams) (list []*system.OrgTreeNode, err error)
    GetSubTree(ctx context.Context, id string, params *model.OrgTreeParams) (*system.OrgTreeNode, error)
    
    // 高级操作
    Move(ctx context.Context, id, newPID string) error
    Copy(ctx context.Context, params *model.OrgCopyParams) (newOrgID string, err error)
    GetStats(ctx context.Context, id string) (*system.OrgStatsInfo, error)
    
    // 批量操作
    BatchDelete(ctx context.Context, ids []string) (*model.BatchResult, error)
    BatchUpdateStatus(ctx context.Context, ids []string, enabled bool) (*model.BatchResult, error)
    
    // 验证
    CheckCodeExists(ctx context.Context, code string) (bool, error)
    CheckCodeExistsExcludeSelf(ctx context.Context, code, id string) (bool, error)
}
```

### 3. 错误处理改进

```go
// 统一错误码定义
const (
    ErrOrgNotFound     = "ORG_NOT_FOUND"
    ErrOrgCodeExists   = "ORG_CODE_EXISTS"
    ErrOrgHasChildren  = "ORG_HAS_CHILDREN"
    ErrOrgHasUsers     = "ORG_HAS_USERS"
    ErrInvalidPID      = "INVALID_PID"
    ErrCircularRef     = "CIRCULAR_REFERENCE"
)

// 统一错误响应
type ErrorResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Details any    `json:"details,omitempty"`
}
```

## 总结

通过以上改进，组织管理API将具备：

1. **更好的RESTful设计**：统一的路径命名和HTTP方法使用
2. **完整的功能覆盖**：支持搜索、分页、批量操作、移动、复制等
3. **更好的用户体验**：提供统计信息、树形结构、路径信息等
4. **更强的扩展性**：支持组织编码、描述、排序等扩展字段
5. **更严格的验证**：参数验证、业务规则检查、错误处理

这些改进将使组织管理模块更加完善和易用。 