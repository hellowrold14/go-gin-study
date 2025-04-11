package common

import "gorm.io/gorm"

type PageResult[T any] struct {
	Total int64 `json:"total"` // 总记录数
	Rows  []*T  `json:"rows"`  // 当前页数据
}

// Paginate 通用分页函数
func Paginate[T any](db *gorm.DB, param *PageParam, condition any) PageResult[T] {
	var records []*T
	var total int64
	offset := (param.Page - 1) * param.PageSize

	// 先查询总记录数
	db.Model(new(T)).Count(&total)

	// 分页查询记录
	db.Where(condition).Offset(offset).Limit(param.PageSize).Find(&records)

	return PageResult[T]{
		Total: total,
		Rows:  records,
	}
}
