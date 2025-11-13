package sysDept

import (
	"IdentifyService/internal/app/system/dao"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sSysDept) CreateDept(ctx context.Context, tx gdb.TX, deptName string, managerID string) (deptID int64, err error) {
	dataInsert := map[string]interface{}{
		dao.SysDept.Columns().ParentId:  0,
		dao.SysDept.Columns().Ancestors: "",
		dao.SysDept.Columns().DeptName:  deptName,
		dao.SysDept.Columns().ManagerID: managerID,
		dao.SysDept.Columns().CreatedAt: gtime.Now(),
		dao.SysDept.Columns().UpdatedAt: gtime.Now(),
	}
	deptID, err = dao.SysDept.Insert(ctx, tx, dataInsert)
	if err != nil {
		return 0, err
	}
	return deptID, nil
}

func (s *sSysDept) IsDeptManager(ctx context.Context, deptID uint64, userID string) (isManager bool, err error) {
	entity, err := dao.SysDept.Get(ctx, deptID)
	if err != nil {
		return false, err
	}
	if entity.ManagerID == userID {
		return true, nil
	}
	return false, nil
}
