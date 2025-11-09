package orm

import (
	"context"
	"database/sql"
)

type dbCommon struct {
	debug bool
	err   error

	//sql   string // 需要执行的 sql 语句
	table string // 表名
	args  []any  // 占位符对应参数
}

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