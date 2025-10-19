package orm

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

type selec[T any] struct {
	dbCommon
	t    *T
	res  []any    // 查询映射结果字段
	cols []string // 查询字段

	join    []string // 查询语法连接 例如：["JOIN ab"]
	on      []string // 查询语法连接条件 例如：["ON a.id = ab.id"]
	where   []string // 查询语法条件 例如：["AND id = ?", "OR account = ?"]
	groupBy []string
	having  string
	orderBy []string
	limit   int64
	offset  int64
}

func (d *selec[T]) setT(t *T) {
	d.t = t
}

func (d *selec[T]) Debug() *T {
	d.debug = true
	return d.t
}

// select a, b from ab where a = 1 group by a order by a limit 1, 2

func (d *selec[T]) FROM(table string) *T {
	if table == "" {
		d.err = errors.New("table name is empty")
	}
	d.table = table
	return d.t
}

func (d *selec[T]) JOIN(table string) *T {
	d.join = append(d.join, "JOIN "+table)
	return d.t
}

func (d *selec[T]) INNER_JOIN(table string) *T {
	d.join = append(d.join, "INNER JOIN "+table)
	return d.t
}

func (d *selec[T]) LEFT_JOIN(table string) *T {
	d.join = append(d.join, "LEFT JOIN "+table)
	return d.t
}

func (d *selec[T]) RIGHT_JOIN(table string) *T {
	d.join = append(d.join, "RIGHT JOIN "+table)
	return d.t
}
func (d *selec[T]) FULL_JOIN(table string) *T {
	d.join = append(d.join, "FULL JOIN "+table)
	return d.t
}

func (d *selec[T]) CROSS_JOIN(table string) *T {
	d.join = append(d.join, "CROSS JOIN "+table)
	return d.t
}

func (d *selec[T]) ON(condition string) *T {
	d.on = append(d.on, condition)
	return d.t
}

func (d *selec[T]) WHERE(where map[string]any) *T {
	if len(where) > 0 {
		d.where = append(d.where, "1 = 1")
	}
	for k, v := range where {
		d.where = append(d.where, k)
		d.args = append(d.args, v)
	}
	return d.t
}

func (d *selec[T]) GROUP_BY(cols ...string) *T {
	d.groupBy = append(d.groupBy, strings.Join(cols, ", "))
	return d.t
}

func (d *selec[T]) HAVING(cond string) *T {
	d.having = cond
	return d.t
}

func (d *selec[T]) ORDER_BY(orders ...string) *T {
	d.orderBy = append(d.orderBy, strings.Join(orders, ", "))
	return d.t
}

func (d *selec[T]) LIMIT(limit int64) *T {
	d.limit = limit
	return d.t
}

func (d *selec[T]) OFFSET(offset int64) *T {
	d.offset = offset
	return d.t
}

func (d *selec[T]) SQL() string {
	var sb strings.Builder
	// 构建查询语句
	sb.WriteString("SELECT " + strings.Join(d.cols, ",") + " FROM " + d.table)
	if len(d.join) > 0 {
		sb.WriteString(" " + strings.Join(d.join, " ")) // JOIN 语句
	}
	if len(d.on) > 0 {
		sb.WriteString(" ON " + strings.Join(d.on, " "))
	}
	if len(d.where) > 0 {
		sb.WriteString(" WHERE " + strings.Join(d.where, " "))
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

func (d *selec[T]) DryRun(ctx context.Context) error {
	if d.err != nil {
		return d.err
	}
	sqlText := d.SQL()
	d.print(ctx, sqlText)
	return nil
}