package admin

import (
	"context"

	"github.com/ChangSZ/gin-boilerplate/internal/pkg/core"
	"github.com/ChangSZ/gin-boilerplate/internal/pkg/password"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql/admin"
)

type CreateAdminData struct {
	Username string // 用户名
	Nickname string // 昵称
	Mobile   string // 手机号
	Password string // 密码
}

func (s *service) Create(ctx context.Context, adminData *CreateAdminData) (id int64, err error) {
	model := admin.NewModel()
	model.Username = adminData.Username
	model.Password = password.GeneratePassword(adminData.Password)
	model.Nickname = adminData.Nickname
	model.Mobile = adminData.Mobile
	model.CreatedUser = core.SessionUserInfo(ctx).UserName
	model.IsUsed = 1
	model.IsDeleted = -1

	id, err = model.Create(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	return
}
