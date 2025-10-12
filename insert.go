package orm

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"testing"
)

type Insert struct {
	dbCommon

	cols []string // 查询字段
	size int      // 插入数据条数
}

// insert into ab (a, b) values (?, ?)

func (d *Insert) insert(rows []Model) {
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

func (d *Insert) INTO(table string) *Insert {
	d.table = table
	return d
}

func (d *Insert) COLUMNS(cols ...string) *Insert {
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

func (d *Insert) VALUES(args ...any) *Insert {
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

func (d *Insert) Exec(ctx context.Context) (int64, error) {
	if _, err := d.Builder(); err != nil {
		return 0, err
	}
	d.debugPrint(ctx)
	res, err := d.db.ExecContext(ctx, d.sql, d.args...)
	if err != nil {
		return 0, fmt.Errorf("db.Exec: %w", err)
	}
	return res.LastInsertId()
}

func (d *Insert) Builder() (string, error) {
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

func (d *Insert) Print(ctx context.Context) error {
	if d.err != nil {
		return d.err
	}
	d.Builder()
	return d.print(ctx)
}

func TestInsert(t *testing.T) {
	// INSERT INTO ab (a, b) VALUES (?, ?)
	err := DB[TestModel](nil).INSERT().INTO("ab").COLUMNS("a", "b").VALUES(1, 2).Print(context.Background())
	if err != nil {
		slog.Error("insert:", "err", err)
	}
}

func TestInsertModel(t *testing.T) {
	// INSERT INTO ab (a, b) VALUES (?, ?)
	err := DB[TestModel](nil).Debug().INSERT(&TestModel{}).INTO("xxx").Print(context.Background())
	if err != nil {
		slog.Error("insert:", "err", err)
	}
}

func TestInsertModels(t *testing.T) {
	// INSERT INTO ab (a, b) VALUES (?, ?), (?, ?)
	err := DB[TestModel](nil).Debug().INSERT(&TestModel{}, &TestModel{}).INTO("xxx").Print(context.Background())
	if err != nil {
		slog.Error("insert:", "err", err)
	}
}