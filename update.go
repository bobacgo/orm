package orm

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"testing"
)

type Update struct {
	dbCommon
	cols  []string
	where []string // 查询语法条件 例如：["AND id = ?", "OR account = ?"]
}

func (d *Update) SET(set map[string]any) *Update {
	if len(set) == 0 {
		d.err = errors.New("empty set")
		return d
	}

	for k, v := range set {
		d.cols = append(d.cols, k)
		d.args = append(d.args, v)
	}
	return d
}

func (d *Update) SET1(row Model) *Update {
	if row == nil {
		d.err = errors.New("empty model")
		return d
	}

	mapping := row.Mapping()
	var (
		columns = make([]string, 0, len(mapping))
		values  = make([]any, 0, len(mapping))
	)
	for _, v := range mapping {
		if v.Column == "id" {
			continue
		}
		//if util.IsZero(v.Value) {
		//	continue
		//}
		columns = append(columns, v.Column+" = ?")
		values = append(values, v.Value)
	}

	if len(columns) == 0 { // 没有更新字段
		return d
	}
	return d
}

func (d *Update) WHERE(where map[string]any) *Update {
	for k, v := range where {
		d.where = append(d.where, k)
		d.args = append(d.args, v)
	}
	return d
}

func (d *Update) Exec(ctx context.Context) (int64, error) {
	if d.err != nil {
		return 0, d.err
	}
	sqlText := "UPDATE " + d.table + " SET " + strings.Join(d.cols, " = ?, ") + " = ? WHERE " + strings.TrimLeft(strings.Join(d.where, " "), "AND ")
	if d.debug {
		slog.InfoContext(ctx, "sql text", "sql", sqlText, "args", d.args)
	}
	res, err := d.db.ExecContext(ctx, sqlText, d.args...)
	if err != nil {
		return 0, fmt.Errorf("db.Exec: %w", err)
	}
	return res.RowsAffected()
}

func TestUpdate(t *testing.T) {
	// UPDATE xx SET a = 1, b = 2 WHERE a > c
	_, err := DB[TestModel](nil).UPDATE("xx").SET(map[string]any{"a": 1, "b": 2}).WHERE(map[string]any{"a": 1}).Exec(context.Background())
	if err != nil {
		slog.Error("update:", "err", err)
	}
}

func TestUpdate1(t *testing.T) {
	// UPDATE xx SET a = 1, b = 2 WHERE a > c
	_, err := DB[TestModel](nil).UPDATE("xx").SET1(&TestModel{}).WHERE(map[string]any{"a": 1}).Exec(context.Background())
	if err != nil {
		slog.Error("update:", "err", err)
	}
}