package openapi

type (
	Pager[T any] struct {
		PageNumber int64 `json:"pageNumber"`
		PageSize   int64 `json:"pageSize"`
		TotalPage  int64 `json:"totalPage"`
		TotalRow   int64 `json:"totalRow"`
		Count      bool  `json:"count"`
		HasNext    bool  `json:"hasNext"`
		Data       []T   `json:"data"`
	}
)

func (p *Pager[T]) FirstRow() int64 {
	return (p.PageNumber - 1) * p.PageSize
}

// Transform takes a transformation function and applies it to each element in the Data slice.
func (p Pager[T]) Transform(transform func(T) interface{}) *Pager[interface{}] {
	newData := make([]interface{}, len(p.Data))
	for i, v := range p.Data {
		newData[i] = transform(v)
	}
	return &Pager[interface{}]{
		PageNumber: p.PageNumber,
		PageSize:   p.PageSize,
		TotalPage:  p.TotalPage,
		TotalRow:   p.TotalRow,
		Count:      p.Count,
		HasNext:    p.HasNext,
		Data:       newData,
	}
}
