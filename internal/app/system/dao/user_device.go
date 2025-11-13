package dao

import (
	"IdentifyService/internal/app/system/model/entity"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type userDeviceDao struct {
	table   string
	group   string
	columns userDeviceColumns
}

type userDeviceColumns struct {
	Id         string
	UserId     string
	DeviceId   string
	DeviceName string
	DeviceIp   string
	LoginType  string
	CreatedAt  string
	UpdatedAt  string
}

var UserDevice = userDeviceDao{
	table: "t_user_device",
	group: "default",
	columns: userDeviceColumns{
		Id:         "id",
		UserId:     "user_id",
		DeviceId:   "device_id",
		DeviceName: "device_name",
		DeviceIp:   "device_ip",
		LoginType:  "login_type",
		CreatedAt:  "created_at",
		UpdatedAt:  "updated_at",
	},
}

func (dao userDeviceDao) Columns() userDeviceColumns {
	return dao.columns
}

func (dao userDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return g.DB(dao.group).Model(dao.table).Ctx(ctx)
}

func (dao userDeviceDao) GetByUserAndDevice(ctx context.Context, userID, deviceID string) (*entity.UserDevice, error) {
	var device entity.UserDevice
	err := dao.Ctx(ctx).
		Where(dao.columns.UserId, userID).
		Where(dao.columns.DeviceId, deviceID).
		Scan(&device)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &device, nil
}

func (dao userDeviceDao) Insert(ctx context.Context, data g.Map) error {
	_, err := dao.Ctx(ctx).Data(data).Insert()
	return err
}

func (dao userDeviceDao) Update(ctx context.Context, id int64, data g.Map) error {
	_, err := dao.Ctx(ctx).Where(dao.columns.Id, id).Data(data).Update()
	return err
}

func (dao userDeviceDao) ListByUser(ctx context.Context, userID string, page, size int) (out []*entity.UserDevice, total int, err error) {
	query := dao.Ctx(ctx).Where(dao.columns.UserId, userID)

	total, err = query.Count()
	if err != nil {
		return nil, 0, err
	}

	err = query.OrderDesc(dao.columns.UpdatedAt).Page(page, size).Scan(&out)

	return out, total, err
}

func (dao userDeviceDao) Delete(ctx context.Context, userID, deviceID string) error {
	_, err := dao.Ctx(ctx).
		Where(dao.columns.UserId, userID).
		Where(dao.columns.DeviceId, deviceID).
		Delete()
	return err
}
