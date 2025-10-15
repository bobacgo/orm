package orm

type dbCommon struct {
	debug bool
	err   error

	sql   string // 需要执行的 sql 语句
	table string // 表名
	args  []any  // 占位符对应参数
}

type Db struct {
	dbCommon
}

func (d *Db) Debug() *Db {
	d.debug = true
	return d
}

type M map[string]any