package menu

import (
	"context"

	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql/menu"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql/menu_action"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx context.Context, menuData *CreateMenuData) (id int64, err error)
	Modify(ctx context.Context, id int64, menuData *UpdateMenuData) (err error)
	List(ctx context.Context, searchData *SearchData) (listData []*menu.Menu, err error)
	UpdateUsed(ctx context.Context, id int64, used int32) (err error)
	UpdateSort(ctx context.Context, id int64, sort int32) (err error)
	Delete(ctx context.Context, id int64) (err error)
	Detail(ctx context.Context, searchOneData *SearchOneData) (info *menu.Menu, err error)

	CreateAction(ctx context.Context, menuActionData *CreateMenuActionData) (id int64, err error)
	ListAction(ctx context.Context, searchListActionData *SearchListActionData) (listData []*menu_action.MenuAction, err error)
	DeleteAction(ctx context.Context, id int64) (err error)
}

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
