package orm

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func (c *dbCommon) print(ctx context.Context) error {
	if c.err != nil {
		return c.err
	}
	sqlText := c.sql
	for _, arg := range c.args {
		switch v := arg.(type) {
		case int, int8, int16, int32, int64:
			sqlText = strings.Replace(sqlText, "?", fmt.Sprintf("%d", v), 1)
		case uint, uint8, uint16, uint32, uint64:
			sqlText = strings.Replace(sqlText, "?", fmt.Sprintf("%d", v), 1)
		case float32, float64:
			sqlText = strings.Replace(sqlText, "?", fmt.Sprintf("%v", v), 1)
		case string:
			sqlText = strings.Replace(sqlText, "?", fmt.Sprintf("'%s'", v), 1)
		case bool:
			sqlText = strings.Replace(sqlText, "?", fmt.Sprintf("%t", v), 1)
		default:
			sqlText = strings.Replace(sqlText, "?", fmt.Sprintf("%v", v), 1)
		}
	}
	slog.InfoContext(ctx, "run sql:", "sql", sqlText)
	return nil
}

func (c *dbCommon) debugPrint(ctx context.Context) {
	if !c.debug {
		return
	}
	slog.InfoContext(ctx, "exec sql:", "sql", c.sql, "args", fmt.Sprintf("%#v", c.args))
}
