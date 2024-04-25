///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package admin_menu

import (
	"fmt"
	"time"

	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *AdminMenu {
	return new(AdminMenu)
}

func NewQueryBuilder() *adminMenuQueryBuilder {
	return new(adminMenuQueryBuilder)
}

func (t *AdminMenu) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type adminMenuQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *adminMenuQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	if qb.limit != 0 {
		ret = ret.Limit(qb.limit)
	}
	ret = ret.Offset(qb.offset)
	return ret
}

func (qb *adminMenuQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&AdminMenu{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *adminMenuQueryBuilder) Update(db *gorm.DB, data *AdminMenu) (cnt int64, err error) {
	db = db.Model(&AdminMenu{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	ret := db.Updates(data)
	err = ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "update err")
	}
	return ret.RowsAffected, nil
}

func (qb *adminMenuQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&AdminMenu{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *adminMenuQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&AdminMenu{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *adminMenuQueryBuilder) First(db *gorm.DB) (*AdminMenu, error) {
	ret := &AdminMenu{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *adminMenuQueryBuilder) QueryOne(db *gorm.DB) (*AdminMenu, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *adminMenuQueryBuilder) QueryAll(db *gorm.DB) ([]*AdminMenu, error) {
	var ret []*AdminMenu
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *adminMenuQueryBuilder) Limit(limit int) *adminMenuQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *adminMenuQueryBuilder) Offset(offset int) *adminMenuQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *adminMenuQueryBuilder) WhereId(p mysql.Predicate, value int64) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) WhereIdIn(value []int64) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) WhereIdNotIn(value []int64) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) OrderById(asc bool) *adminMenuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *adminMenuQueryBuilder) WhereAdminId(p mysql.Predicate, value int64) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "admin_id", p),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) WhereAdminIdIn(value []int64) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "admin_id", "IN"),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) WhereAdminIdNotIn(value []int64) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "admin_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) OrderByAdminId(asc bool) *adminMenuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "admin_id "+order)
	return qb
}

func (qb *adminMenuQueryBuilder) WhereMenuId(p mysql.Predicate, value int64) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "menu_id", p),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) WhereMenuIdIn(value []int64) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "menu_id", "IN"),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) WhereMenuIdNotIn(value []int64) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "menu_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) OrderByMenuId(asc bool) *adminMenuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "menu_id "+order)
	return qb
}

func (qb *adminMenuQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) WhereCreatedAtIn(value []time.Time) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) OrderByCreatedAt(asc bool) *adminMenuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *adminMenuQueryBuilder) WhereCreatedUser(p mysql.Predicate, value string) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) WhereCreatedUserIn(value []string) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) WhereCreatedUserNotIn(value []string) *adminMenuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *adminMenuQueryBuilder) OrderByCreatedUser(asc bool) *adminMenuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}