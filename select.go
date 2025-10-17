package orm

import (
	"errors"
	"strconv"
	"strings"
)

type selec[T any] struct {
	dbCommon
	t   *T
	res []any // 查询映射结果字段
}

func (d *selec[T]) setT(t *T) {
	d.t = t
}

// select a, b from ab where a = 1 group by a order by a limit 1, 2

func (d *selec[T]) FROM(table string) *T {
	if table == "" {
		d.err = errors.New("table name is empty")
	}
	d.table = table
	d.sql += " FROM " + d.table
	return d.t
}

func (d *selec[T]) JOIN(table string) *T {
	d.sql += " JOIN " + table
	return d.t
}

func (d *selec[T]) INNER_JOIN(table string) *T {
	d.sql += " INNER JOIN " + table
	return d.t
}

func (d *selec[T]) LEFT_JOIN(table string) *T {
	d.sql += " LEFT JOIN " + table
	return d.t
}

func (d *selec[T]) RIGHT_JOIN(table string) *T {
	d.sql += " RIGHT JOIN " + table
	return d.t
}
func (d *selec[T]) FULL_JOIN(table string) *T {
	d.sql += " FULL JOIN " + table
	return d.t
}

func (d *selec[T]) CROSS_JOIN(table string) *T {
	d.sql += " CROSS JOIN " + table
	return d.t
}

func (d *selec[T]) ON(condition string) *T {
	d.sql += " ON " + condition
	return d.t
}

func (d *selec[T]) WHERE(where map[string]any) *T {
	if len(where) > 0 {
		d.sql += " WHERE " // TODO bug
		for k, v := range where {
			d.sql += k
			d.args = append(d.args, v)
		}
	}
	return d.t
}

func (d *selec[T]) GROUP_BY(cols ...string) *T {
	d.sql += " GROUP BY " + strings.Join(cols, ", ")
	return d.t
}

func (d *selec[T]) HAVING(text string) *T {
	d.sql += " HAVING " + text
	return d.t
}

func (d *selec[T]) ORDER_BY(orders ...string) *T {
	d.sql += " ORDER BY " + strings.Join(orders, ", ")
	return d.t
}

func (d *selec[T]) LIMIT(limit int64) *T {
	d.sql += " LIMIT " + strconv.FormatInt(limit, 10)
	return d.t
}

func (d *selec[T]) OFFSET(offset int64) *T {
	d.sql += " OFFSET " + strconv.FormatInt(offset, 10)
	return d.t
}