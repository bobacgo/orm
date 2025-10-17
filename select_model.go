package orm

import (
	"context"
	"database/sql"
	"log/slog"
	"strings"
)

type SelectModel struct {
	*selec[SelectModel]
}

func SELECT(row Model) *SelectModel {
	s := &SelectModel{
		&selec[SelectModel]{
			dbCommon: dbCommon{}, // TODO
		},
	}
	var cols []string
	for _, v := range row.Mapping() {
		cols = append(cols, v.Column)
		s.res = append(s.res, v.Result)
	}
	s.setT(s)
	s.sql = "SELECT " + strings.Join(cols, ", ")
	return s
}

func (d *SelectModel) Query(ctx context.Context, db *sql.DB) error {
	if d.err != nil {
		return d.err
	}
	sqlText := d.sql
	if d.debug {
		slog.InfoContext(ctx, "sql text", "sql", sqlText, "args", d.args)
	}
	return db.QueryRowContext(ctx, sqlText, d.args...).Scan(d.res...)
}