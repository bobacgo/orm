package orm

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Delete struct {
	dbCommon
	where []string // 查询条件
}

func DELETE() *Delete {
	del := &Delete{
		dbCommon: dbCommon{},
	}
	return del
}

func (d *Delete) FROM(table string) *Delete {
	// 检查 table 是否为空
	if table == "" {
		d.err = fmt.Errorf("table name is empty")
	}
	d.table = table
	return d
}

func (d *Delete) WHERE(where map[string]any) *Delete {
	for k, v := range where {
		d.where = append(d.where, k)
		d.args = append(d.args, v)
	}
	return d
}

func (d *Delete) SQL() error {
	if d.err != nil {
		return d.err
	}
	d.sql = "DELETE FROM " + d.table + " WHERE " + strings.TrimLeft(strings.Join(d.where, " "), "AND ")
	return nil
}

func (d *Delete) Exec(ctx context.Context, db *sql.DB) (int64, error) {
	if err := d.SQL(); err != nil {
		return 0, err
	}
	d.debugPrint(ctx)
	res, err := db.ExecContext(ctx, d.sql, d.args...)
	if err != nil {
		return 0, fmt.Errorf("db..Exec: %w", err)
	}
	return res.RowsAffected()
}