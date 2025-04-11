package common

type PageParam struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

func NewPageParam() *PageParam {
	return &PageParam{
		Page:     1,
		PageSize: 10,
	}
}
