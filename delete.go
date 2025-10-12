package orm

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"testing"
)

type Delete struct {
	dbCommon
	where []string // 查询条件
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

func (d *Delete) Exec(ctx context.Context) (int64, error) {
	if err := d.SQL(); err != nil {
		return 0, err
	}
	d.debugPrint(ctx)
	res, err := d.db.ExecContext(ctx, d.sql, d.args...)
	if err != nil {
		return 0, fmt.Errorf("db..Exec: %w", err)
	}
	return res.RowsAffected()
}

func TestDelete(t *testing.T) {
	// DELETE FROM xxx WHERE a = 1
	_, err := DB[TestModel](nil).Debug().DELETE().FROM("xxx").WHERE(map[string]any{"a": 1}).Exec(context.Background())
	if err != nil {
		slog.Error("delete:", "err", err)
	}
}