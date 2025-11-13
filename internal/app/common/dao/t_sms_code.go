package dao

import (
	"IdentifyService/internal/app/common/dao/internal"
	"IdentifyService/internal/app/common/model"
	"IdentifyService/internal/app/common/model/entity"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type TSmsCodeDao struct {
	*internal.TSmsCodeDao
}

var (
	TSmsCode = TSmsCodeDao{
		internal.NewTSmsCodeDao(),
	}
)

func (d *TSmsCodeDao) AddSmsCode(ctx context.Context, data map[string]interface{}) (smsCode *entity.TSmsCode, err error) {
	_, err = d.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (d *TSmsCodeDao) GetLastSmsCode(ctx context.Context, phone string, bussinessType model.SMSBusinessType, code string) (smsCode *entity.TSmsCode, err error) {
	smsCode = &entity.TSmsCode{}

	err = d.Ctx(ctx).
		Where(d.Columns().Phone, phone).
		Where(d.Columns().BusinessType, bussinessType).
		Where(d.Columns().Code, code).
		Scan(&smsCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("验证码不存在")
		}
		return nil, err
	}

	return smsCode, nil
}

func (d *TSmsCodeDao) UpdateSmsCodeStatus(ctx context.Context, id int64, status model.SMSCodeStatus) (err error) {
	_, err = d.Ctx(ctx).Where(d.Columns().Id, id).Data(map[string]interface{}{
		d.Columns().Status:    status,
		d.Columns().UpdatedAt: time.Now().Unix(),
	}).Update()
	return err
}
