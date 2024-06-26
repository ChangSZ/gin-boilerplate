package admin

import (
	"context"

	"github.com/ChangSZ/gin-boilerplate/configs"
	"github.com/ChangSZ/gin-boilerplate/internal/pkg/core"
	"github.com/ChangSZ/gin-boilerplate/internal/pkg/password"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql/admin"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/redis"
)

func (s *service) ResetPassword(ctx context.Context, id int64) (err error) {
	data := map[string]interface{}{
		"password":     password.ResetPassword(),
		"updated_user": core.SessionUserInfo(ctx).UserName,
	}

	qb := admin.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return err
	}

	redis.Cache().Del(ctx, configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id))
	return
}
