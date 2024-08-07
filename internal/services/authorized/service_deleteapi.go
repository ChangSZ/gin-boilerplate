package authorized

import (
	"context"

	"gorm.io/gorm"

	"github.com/ChangSZ/gin-boilerplate/configs"
	"github.com/ChangSZ/gin-boilerplate/internal/pkg/core"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql/authorized_api"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/redis"
)

func (s *service) DeleteAPI(ctx context.Context, id int64) (err error) {
	// 先查询 id 是否存在
	authorizedApiInfo, err := authorized_api.NewQueryBuilder().
		WhereIsDeleted(mysql.EqualPredicate, -1).
		WhereId(mysql.EqualPredicate, id).
		First(mysql.DB().GetDbR().WithContext(ctx))

	if err == gorm.ErrRecordNotFound {
		return nil
	}

	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": core.SessionUserInfo(ctx).UserName,
	}

	qb := authorized_api.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return err
	}

	redis.Cache().Del(ctx, configs.RedisKeyPrefixSignature+authorizedApiInfo.BusinessKey)
	return
}
