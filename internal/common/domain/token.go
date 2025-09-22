package domain

import (
	"IdentifyService/internal/common/dao"
	"context"
	"strings"
	"sync"
)

var (
	tokenOnce     sync.Once
	tokenInstance *token
)

type token struct {
}

func NewToken() *token {
	tokenOnce.Do(func() {
		tokenInstance = &token{}
	})
	return tokenInstance
}

func (t *token) AddToBlacklist(ctx context.Context, jti string, operatorID string) (err error) {
	_, err = dao.TokenBlacklist.Ctx(ctx).Insert(map[string]interface{}{
		dao.TokenBlacklist.Columns().ID:         jti,
		dao.TokenBlacklist.Columns().OperatorID: operatorID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil
		}
	}
	return
}

func (t *token) InBlacklist(ctx context.Context, jti string) (exists bool, err error) {
	exists, err = dao.TokenBlacklist.Ctx(ctx).Where(dao.TokenBlacklist.Columns().ID, jti).Exist()

	return
}
