package request

type CommonPage struct {
	Page int `json:"page"` // 页码
	Size int `json:"size"` // 页大小
}

func (c *CommonPage) Offset() int {
	return (c.Page - 1) * c.Size
}
