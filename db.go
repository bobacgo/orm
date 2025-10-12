package orm

import (
	"database/sql"
)

type dbCommon struct {
	db    *sql.DB
	debug bool
	err   error

	sql   string // 需要执行的 sql 语句
	table string // 表名
	args  []any  // 占位符对应参数
}

type Db[T any, P PModel[T]] struct {
	dbCommon

	slect  *Select[T, P]
	insert *Insert
	update *Update
	delete *Delete
}

func DB[T any, P PModel[T]](db *sql.DB) *Db[T, P] {
	return &Db[T, P]{
		dbCommon: dbCommon{
			db: db,
		},
	}
}

func (d *Db[T, P]) Debug() *Db[T, P] {
	d.debug = true
	return d
}

func (d *Db[T, P]) SELECT(row Model) *Select[T, P] {
	d.slect = &Select[T, P]{
		dbCommon: d.dbCommon,
	}
	for _, v := range row.Mapping() {
		d.slect.cols = append(d.slect.cols, v.Column)
		d.slect.res = append(d.slect.res, v.Result)
	}
	return d.slect
}

func (d *Db[T, P]) SELECT1(rows *[]*T) *Select[T, P] {
	d.slect = &Select[T, P]{
		dbCommon: d.dbCommon,
		row:      new(T),
		rows:     rows,
	}
	mappings := P(d.slect.row).Mapping()
	for _, v := range mappings {
		d.slect.cols = append(d.slect.cols, v.Column)
		d.slect.res = append(d.slect.res, v.Result)
	}
	return d.slect
}

func (d *Db[T, P]) SELECT2(cols ...string) *Select[T, P] {
	d.slect = &Select[T, P]{
		dbCommon: d.dbCommon,
		cols:     cols,
	}
	return d.slect
}

func (d *Db[T, P]) INSERT(rows ...Model) *Insert {
	d.insert = &Insert{
		dbCommon: d.dbCommon,
	}
	d.insert.insert(rows)
	return d.insert
}

func (d *Db[T, P]) UPDATE(tableName string) *Update {
	d.update = &Update{
		dbCommon: d.dbCommon,
	}
	d.dbCommon.table = tableName
	return d.update
}

func (d *Db[T, P]) DELETE() *Delete {
	d.delete = &Delete{
		dbCommon: d.dbCommon,
	}
	return d.delete
}

type M map[string]any