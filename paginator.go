package orm

import (
	"gorm.io/gorm"
)

// Paginator 分页参数
type Paginator struct {
	DB       *gorm.DB
	Page     int
	PageSize int
	OrderBy  []string
	ShowSQL  bool
}

// Data 分页返回
type Data struct {
	Count    int64       `json:"count"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Rows     interface{} `json:"results"`
}

func (p *Paginator) Paginate(result interface{}) *Data {
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	var paginator Data
	var count int64
	var offset int

	db.Model(result).Count(&count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.PageSize
	}

	db.Limit(p.PageSize).Offset(offset).Find(result)

	paginator.Count = count
	paginator.Rows = result
	paginator.Page = p.Page
	paginator.PageSize = p.PageSize
	return &paginator
}

