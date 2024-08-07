package authorized

import (
	"context"

	"gorm.io/gorm"

	"github.com/ChangSZ/gin-boilerplate/configs"
	"github.com/ChangSZ/gin-boilerplate/internal/pkg/core"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql/authorized"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/redis"
)

func (s *service) UpdateUsed(ctx context.Context, id int64, used int32) (err error) {
	authorizedInfo, err := authorized.NewQueryBuilder().
		WhereIsDeleted(mysql.EqualPredicate, -1).
		WhereId(mysql.EqualPredicate, id).
		First(mysql.DB().GetDbR().WithContext(ctx))

	if err == gorm.ErrRecordNotFound {
		return nil
	}

	data := map[string]interface{}{
		"is_used":      used,
		"updated_user": core.SessionUserInfo(ctx).UserName,
	}

	qb := authorized.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return err
	}

	redis.Cache().Del(ctx, configs.RedisKeyPrefixSignature+authorizedInfo.BusinessKey)
	return
}
