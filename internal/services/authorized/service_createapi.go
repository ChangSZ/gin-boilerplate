package authorized

import (
	"context"

	"github.com/ChangSZ/gin-boilerplate/configs"
	"github.com/ChangSZ/gin-boilerplate/internal/pkg/core"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql/authorized_api"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/redis"
)

type CreateAuthorizedAPIData struct {
	BusinessKey string `json:"business_key"` // 调用方key
	Method      string `json:"method"`       // 请求方法
	API         string `json:"api"`          // 请求地址
}

func (s *service) CreateAPI(ctx context.Context, authorizedAPIData *CreateAuthorizedAPIData) (id int64, err error) {
	model := authorized_api.NewModel()
	model.BusinessKey = authorizedAPIData.BusinessKey
	model.Method = authorizedAPIData.Method
	model.Api = authorizedAPIData.API
	model.CreatedUser = core.SessionUserInfo(ctx).UserName
	model.IsDeleted = -1

	id, err = model.Create(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}

	redis.Cache().Del(ctx, configs.RedisKeyPrefixSignature+authorizedAPIData.BusinessKey)
	return
}
