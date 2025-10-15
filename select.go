package orm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

type Select[T any, P PModel[T]] struct {
	dbCommon
	cols []string // 查询字段
	res  []any    // 查询映射结果字段

	row  *T
	rows *[]*T

	where   []string // 查询语法条件 例如：["AND id = ?", "OR account = ?"]
	groupBy []string
	having  string
	orderBy []string
	limit   int64
	offset  int64
}

// select a, b from ab where a = 1 group by a order by a limit 1, 2

func (d *Select[T, P]) FROM(table string) *Select[T, P] {
	if table == "" {
		d.err = errors.New("table name is empty")
	}
	d.table = table
	return d
}

func (d *Select[T, P]) WHERE(where map[string]any) *Select[T, P] {
	for k, v := range where {
		d.where = append(d.where, k)
		d.args = append(d.args, v)
	}
	return d
}

func (d *Select[T, P]) GROUP_BY(cols ...string) *Select[T, P] {
	d.groupBy = cols
	return d
}

func (d *Select[T, P]) HAVING(text string) *Select[T, P] {
	d.having = text
	return d
}

func (d *Select[T, P]) ORDER_BY(orders ...string) *Select[T, P] {
	d.orderBy = orders
	return d
}

func (d *Select[T, P]) LIMIT(limit int64) *Select[T, P] {
	d.limit = limit
	return d
}

func (d *Select[T, P]) OFFSET(offset int64) *Select[T, P] {
	d.offset = offset
	return d
}

// 查到单个
func (d *Select[T, P]) QueryRow(ctx context.Context, db *sql.DB) error {
	if d.err != nil {
		return d.err
	}
	sqlText := d.SQL()
	if d.debug {
		slog.InfoContext(ctx, "sql text", "sql", sqlText, "args", d.args)
	}
	return db.QueryRowContext(ctx, sqlText, d.args...).Scan(d.res...)
}

// 查询多个
func (d *Select[T, P]) Query(ctx context.Context, db *sql.DB) error {
	if d.err != nil {
		return d.err
	}

	sqlText := d.SQL()
	if d.debug {
		slog.InfoContext(ctx, "sql text", "sql", sqlText, "args", d.args)
	}
	rows, err := db.QueryContext(ctx, sqlText, d.args...)
	if err != nil {
		return fmt.Errorf("stmt.Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(d.res...); err != nil {
			return fmt.Errorf("rows.Scan: %w", err)
		}
		cp := *d.row
		*d.rows = append(*d.rows, &cp)
	}
	return rows.Err()
}

func (d *Select[T, P]) SQL() string {
	var sb strings.Builder
	// 构建查询语句
	sb.WriteString("SELECT " + strings.Join(d.cols, ",") + " FROM " + d.table)

	if len(d.where) > 0 {
		sb.WriteString(" WHERE " + strings.TrimLeft(strings.Join(d.where, " "), "AND "))
	}
	if len(d.groupBy) > 0 {
		sb.WriteString(" GROUP BY " + strings.Join(d.groupBy, ", "))
	}
	if d.having != "" {
		sb.WriteString(" HAVING " + d.having)
	}
	if len(d.orderBy) > 0 {
		sb.WriteString(" ORDER BY " + strings.Join(d.orderBy, ", "))
	}
	if d.limit > 0 {
		sb.WriteString(" LIMIT " + strconv.FormatInt(d.limit, 10))
	}
	if d.offset > 0 {
		sb.WriteString(" OFFSET " + strconv.FormatInt(d.offset, 10))
	}

	return sb.String()
}