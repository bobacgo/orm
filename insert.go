package orm

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Insert[T Model] struct {
	dbCommon

	cols []string // 查询字段
	size int      // 插入数据条数
}

func INSERT[T Model](rows ...T) *Insert[T] {
	insert := &Insert[T]{
		dbCommon: dbCommon{},
	}
	insert.insert(rows)
	return insert
}

func INSERT1() *Insert[*emptyModel] {
	return INSERT[*emptyModel]()
}

// insert into ab (a, b) values (?, ?)

func (d *Insert[T]) insert(rows []T) {
	switch len(rows) {
	case 1:
		d.size = 1
		mapping := rows[0].Mapping()
		for _, v := range mapping {
			if v.Column == "id" {
				continue
			}
			d.cols = append(d.cols, v.Column)
			d.args = append(d.args, v.Value)
		}
	default:
		d.size = len(rows)
		for i, row := range rows {
			mapping := row.Mapping()
			for _, v := range mapping {
				if v.Column == "id" {
					continue
				}
				if i == 0 {
					d.cols = append(d.cols, v.Column)
				}
				d.args = append(d.args, v.Value)
			}
		}
	}
}

func (d *Insert[T]) INTO(table string) *Insert[T] {
	d.table = table
	return d
}

func (d *Insert[T]) COLUMNS(cols ...string) *Insert[T] {
	if d.size != 0 {
		d.err = fmt.Errorf("columns and models can only be used once")
		return d
	}
	if len(cols) == 0 {
		d.err = fmt.Errorf("columns is empty")
		return d
	}
	d.cols = cols
	return d
}

func (d *Insert[T]) VALUES(args ...any) *Insert[T] {
	if d.size != 0 {
		d.err = fmt.Errorf("columns and models can only be used once")
		return d
	}
	if len(args) == 0 {
		d.err = fmt.Errorf("args is empty")
		return d
	}
	d.args = append(d.args, args...)
	d.size = 1
	return d
}

func (d *Insert[T]) Exec(ctx context.Context, db *sql.DB) (int64, error) {
	if _, err := d.SQL(); err != nil {
		return 0, err
	}
	d.debugPrint(ctx)
	res, err := db.ExecContext(ctx, d.sql, d.args...)
	if err != nil {
		return 0, fmt.Errorf("db.Exec: %w", err)
	}
	return res.LastInsertId()
}

func (d *Insert[T]) SQL() (string, error) {
	if d.err != nil {
		return "", d.err
	}
	var builder strings.Builder
	builder.WriteString("INSERT INTO " + d.table + " (" + strings.Join(d.cols, ", ") + ") VALUES ")
	for i := range d.size {
		builder.WriteString("(" + strings.Repeat("?, ", len(d.cols)-1) + "?)")
		if i < d.size-1 {
			builder.WriteString(", ")
		}
	}
	d.sql = builder.String()
	return d.sql, nil
}

func (d *Insert[T]) Print(ctx context.Context) error {
	if d.err != nil {
		return d.err
	}
	d.SQL()
	return d.print(ctx)
}