package orm

import "github.com/jinzhu/gorm"

type Page struct {
	PageNo   int `json:"page_no"`
	PageSize int `json:"page_size"`
}

func (p *Page) Apply(db *gorm.DB) {
	// page size 大小由业务代码自己决定
	*db = *db.Offset((p.PageNo - 1) * p.PageSize)
	*db = *db.Limit(p.PageSize)
}
