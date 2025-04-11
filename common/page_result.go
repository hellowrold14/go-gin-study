package common

type PageResult[T any] struct {
	Total int64 `json:"total"` // 总记录数
	Rows  []T   `json:"rows"`  // 当前页数据
}
