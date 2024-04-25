///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package menu_action

import (
	"fmt"
	"time"

	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *MenuAction {
	return new(MenuAction)
}

func NewQueryBuilder() *menuActionQueryBuilder {
	return new(menuActionQueryBuilder)
}

func (t *MenuAction) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type menuActionQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *menuActionQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *menuActionQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&MenuAction{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *menuActionQueryBuilder) Update(db *gorm.DB, data *MenuAction) (cnt int64, err error) {
	db = db.Model(&MenuAction{})

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

func (qb *menuActionQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&MenuAction{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *menuActionQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&MenuAction{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *menuActionQueryBuilder) First(db *gorm.DB) (*MenuAction, error) {
	ret := &MenuAction{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *menuActionQueryBuilder) QueryOne(db *gorm.DB) (*MenuAction, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *menuActionQueryBuilder) QueryAll(db *gorm.DB) ([]*MenuAction, error) {
	var ret []*MenuAction
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *menuActionQueryBuilder) Limit(limit int) *menuActionQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *menuActionQueryBuilder) Offset(offset int) *menuActionQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *menuActionQueryBuilder) WhereId(p mysql.Predicate, value int64) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereIdIn(value []int64) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereIdNotIn(value []int64) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) OrderById(asc bool) *menuActionQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *menuActionQueryBuilder) WhereMenuId(p mysql.Predicate, value int64) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "menu_id", p),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereMenuIdIn(value []int64) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "menu_id", "IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereMenuIdNotIn(value []int64) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "menu_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) OrderByMenuId(asc bool) *menuActionQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "menu_id "+order)
	return qb
}

func (qb *menuActionQueryBuilder) WhereMethod(p mysql.Predicate, value string) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "method", p),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereMethodIn(value []string) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "method", "IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereMethodNotIn(value []string) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "method", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) OrderByMethod(asc bool) *menuActionQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "method "+order)
	return qb
}

func (qb *menuActionQueryBuilder) WhereApi(p mysql.Predicate, value string) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "api", p),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereApiIn(value []string) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "api", "IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereApiNotIn(value []string) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "api", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) OrderByApi(asc bool) *menuActionQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "api "+order)
	return qb
}

func (qb *menuActionQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereIsDeletedIn(value []int32) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereIsDeletedNotIn(value []int32) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) OrderByIsDeleted(asc bool) *menuActionQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *menuActionQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereCreatedAtIn(value []time.Time) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) OrderByCreatedAt(asc bool) *menuActionQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *menuActionQueryBuilder) WhereCreatedUser(p mysql.Predicate, value string) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereCreatedUserIn(value []string) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereCreatedUserNotIn(value []string) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) OrderByCreatedUser(asc bool) *menuActionQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *menuActionQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereUpdatedAtIn(value []time.Time) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) OrderByUpdatedAt(asc bool) *menuActionQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *menuActionQueryBuilder) WhereUpdatedUser(p mysql.Predicate, value string) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereUpdatedUserIn(value []string) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) WhereUpdatedUserNotIn(value []string) *menuActionQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuActionQueryBuilder) OrderByUpdatedUser(asc bool) *menuActionQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}
