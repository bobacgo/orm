package orm

import (
	"context"
	"database/sql"
	"maps"
	"slices"
)

// sql.DB | sql.Tx
type Execer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type M map[string]any
type Mapping struct {
	Column string
	Result any // query result (pointer)
	Value  any // insert, update value
}

type dbCommon struct {
	debug bool
	err   error

	//sql   string // 需要执行的 sql 语句
	table string // 表名
	args  []any  // 占位符对应参数
}

func (c *dbCommon) excludeNil(m M) ([]string, []any) {
	cds, vs := make([]string, 0, len(m)), make([]any, 0)
	keys := slices.Sorted(maps.Keys(m))
	for _, k := range keys { // 按键排序
		cds = append(cds, k)
		if v := m[k]; v != nil { // 没有占位符
			vs = append(vs, m[k])
		}
	}
	return cds, vs
}

func (d *dbCommon) omitCol(omitcols, cols []string, val []any) ([]string, []any) {
	// 如果后面调用 Omit 方法，忽略更新字段
	newcols, newval := make([]string, 0), make([]any, 0)
	for i, cl := range cols {
		if slices.Contains(omitcols, cl) {
			continue
		}
		newcols = append(newcols, cl)
		newval = append(newval, val[i])
	}
	return newcols, newval
}
