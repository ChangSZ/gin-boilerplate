package menu

import (
	"context"

	"github.com/ChangSZ/gin-boilerplate/internal/pkg/core"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql/menu"
)

func (s *service) Delete(ctx context.Context, id int64) (err error) {
	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": core.SessionUserInfo(ctx).UserName,
	}

	qb := menu.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return err
	}

	return
}
