package interfaces

import (
	"context"

	"IdentifyService/internal/system/model"

	"IdentifyService/api/v1/system"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IRole interface {
		// 添加角色
		Create(ctx context.Context, tx gdb.TX, req *system.RoleCreateReq) (res *system.RoleCreateRes, err error)
		// 根据角色ID列表批量删除
		DeleteByIDs(ctx context.Context, tx gdb.TX, ids []int64) (err error)
		// 根据ID编辑角色
		Edit(ctx context.Context, tx gdb.TX, req *system.RoleEditReq) (err error)
		// 根据角色ID获取角色详细信息
		Get(ctx context.Context, id int64) (out *model.Role, err error)
		// 根据角色名称获取角色详细信息
		GetByName(ctx context.Context, name string) (out *model.Role, err error)

		// 根据角色ID获取角色树形结构
		GetTreeByRoleID(ctx context.Context, id int64) (out *system.RoleNode, err error)
	}
)
