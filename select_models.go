package orm

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
)

type SelectModels[T any, P PModel[T]] struct {
	*selec[SelectModels[T, P]]
	row  *T
	rows *[]*T
}

func SELECT1[T any, P PModel[T]](rows *[]*T) *SelectModels[T, P] {
	s := &SelectModels[T, P]{
		selec: &selec[SelectModels[T, P]]{
			dbCommon: dbCommon{}, // TODO
		},
		row:  new(T),
		rows: rows,
	}
	var cols []string
	mappings := P(s.row).Mapping()
	for _, v := range mappings {
		cols = append(cols, v.Column)
		s.res = append(s.res, v.Result)
	}
	s.setT(s)
	s.sql = "SELECT " + strings.Join(cols, ", ")
	return s
}

func (d *SelectModels[T, P]) Query(ctx context.Context, db *sql.DB) error {
	if d.err != nil {
		return d.err
	}

	sqlText := d.sql
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