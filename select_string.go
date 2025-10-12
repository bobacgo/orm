package orm

import (
	"fmt"
	"testing"
)

type SelectString struct {
	*Select[emptyModel, *emptyModel]
}

func SELECT2(cols ...string) SelectString {
	return SelectString{
		&Select[emptyModel, *emptyModel]{
			dbCommon: dbCommon{},
			cols:     cols,
		},
	}
}

func TestSelectString(t *testing.T) {
	// SELECT a, b FROM xx WHERE id = 1 GROUP BY a HAVING a > 0 ORDER BY a desc, b LIMIT 1, 1
	sqlText := SELECT2("a", "b").
		FROM("xx").
		WHERE(map[string]any{"AND id = ?": 1}).
		GROUP_BY("a").
		HAVING("a > 0").
		ORDER_BY("a desc", "b").
		OFFSET(1).LIMIT(1).
		SQL()
	fmt.Println(sqlText)
}