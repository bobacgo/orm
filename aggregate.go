package orm

type count struct {
	v   *int64
	col string
}

func COUNT(col string, v *int64) *count {
	return &count{
		v:   v,
		col: "COUNT(" + col + ")",
	}
}

func (c *count) TableName() string {
	return ""
}

func (c *count) Mapping(_ bool) map[string]any {
	return map[string]any{
		c.col: c.v,
	}
}