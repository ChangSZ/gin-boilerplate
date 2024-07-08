package pkg

import "text/template"

// Make sure that the template compiles during package initialization
func parseTemplateOrPanic(t string) *template.Template {
	tpl, err := template.New("output_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}

var outputTemplate = parseTemplateOrPanic(`
///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package {{.PkgName}}

import (
	"fmt"
	"time"

	"github.com/ChangSZ/gin-boilerplate/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *{{.StructName}} {
	return new({{.StructName}})
}

func NewQueryBuilder() *{{.QueryBuilderName}} {
	return new({{.QueryBuilderName}})
}

func (t *{{.StructName}}) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type {{.QueryBuilderName}} struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *{{.QueryBuilderName}}) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *{{.QueryBuilderName}}) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&{{.StructName}}{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	ret := db.Updates(m)
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "updates err")
	}
	return ret.RowsAffected, nil
}


func (qb *{{.QueryBuilderName}}) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret :=db.Delete(&{{.StructName}}{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *{{.QueryBuilderName}}) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&{{.StructName}}{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *{{.QueryBuilderName}}) First(db *gorm.DB) (*{{.StructName}}, error) {
	ret := &{{.StructName}}{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *{{.QueryBuilderName}}) QueryOne(db *gorm.DB) (*{{.StructName}}, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *{{.QueryBuilderName}}) QueryAll(db *gorm.DB) ([]*{{.StructName}}, error) {
	var ret []*{{.StructName}}
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *{{.QueryBuilderName}}) Limit(limit int) *{{.QueryBuilderName}} {
	qb.limit = limit
	return qb
}

func (qb *{{.QueryBuilderName}}) Offset(offset int) *{{.QueryBuilderName}} {
	qb.offset = offset
	return qb
}

{{$queryBuilderName := .QueryBuilderName}}
{{range .OptionFields}}
func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}(p mysql.Predicate, value {{.FieldType}}) *{{$queryBuilderName}} {
	 qb.where = append(qb.where, struct {
		prefix string
		value interface{}
	}{
		fmt.Sprintf("%v %v ?", "{{.ColumnName}}", p),
		value,
	})
	return qb
}

func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}In(value []{{.FieldType}}) *{{$queryBuilderName}} {
	 qb.where = append(qb.where, struct {
		prefix string
		value interface{}
	}{
		fmt.Sprintf("%v %v ?", "{{.ColumnName}}", "IN"),
		value,
	})
	return qb
}

func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}NotIn(value []{{.FieldType}}) *{{$queryBuilderName}} {
	 qb.where = append(qb.where, struct {
		prefix string
		value interface{}
	}{
		fmt.Sprintf("%v %v ?", "{{.ColumnName}}", "NOT IN"),
		value,
	})
	return qb
}

func (qb *{{$queryBuilderName}}) OrderBy{{call $.Helpers.Titelize .FieldName}}(asc bool) *{{$queryBuilderName}} {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "{{.ColumnName}} " + order)
	return qb
}
{{end}}
`)
