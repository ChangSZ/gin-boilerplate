package admin

import (
	"context"

	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql/admin"
)

type SearchData struct {
	Page     int    // 第几页
	PageSize int    // 每页显示条数
	Username string // 用户名
	Nickname string // 昵称
	Mobile   string // 手机号
}

func (s *service) PageList(ctx context.Context, searchData *SearchData) (listData []*admin.Admin, err error) {

	page := searchData.Page
	if page == 0 {
		page = 1
	}

	pageSize := searchData.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	qb := admin.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	if searchData.Username != "" {
		qb.WhereUsername(mysql.EqualPredicate, searchData.Username)
	}

	if searchData.Nickname != "" {
		qb.WhereNickname(mysql.EqualPredicate, searchData.Nickname)
	}

	if searchData.Mobile != "" {
		qb.WhereMobile(mysql.EqualPredicate, searchData.Mobile)
	}

	listData, err = qb.
		Limit(pageSize).
		Offset(offset).
		OrderById(false).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	return
}
