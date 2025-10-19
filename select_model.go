package orm

import (
	"context"
	"database/sql"
)

type SelectModel struct {
	*selec[SelectModel]
}

func SELECT(row Model) *SelectModel {
	s := &SelectModel{
		&selec[SelectModel]{
			dbCommon: dbCommon{},
		},
	}
	for _, v := range row.Mapping() {
		s.cols = append(s.cols, v.Column)
		s.res = append(s.res, v.Result)
	}
	s.setT(s)
	return s
}

func (d *SelectModel) Query(ctx context.Context, db *sql.DB) error {
	if d.err != nil {
		return d.err
	}
	sqlText := d.SQL()
	d.debugPrint(ctx, sqlText)
	return db.QueryRowContext(ctx, sqlText, d.args...).Scan(d.res...)
}