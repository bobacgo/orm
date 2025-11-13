package orm

import (
	"context"
	"database/sql"
	"strings"
)

type dbCommon struct {
	debug bool
	err   error

	//sql   string // 需要执行的 sql 语句
	table string // 表名
	args  []any  // 占位符对应参数
}

// sql.DB | sql.Tx
type Execer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type M map[string]any
type Mapping struct {
	Column string
	Result any // query result (pointer)
	Value  any // insert, update value
}

// expandInClause expands "in (?)" placeholders when the value is a comma-separated string
// For example: "id in (?)" with value "1,2,3" becomes "id in (?,?,?)" with values [1,2,3]
func expandInClause(clause string, value any) (string, []any) {
	// Check if the clause contains "in (?)"
	lowerClause := strings.ToLower(clause)
	if !strings.Contains(lowerClause, "in (?)") {
		return clause, []any{value}
	}
	
	// Check if the value is a string
	strValue, ok := value.(string)
	if !ok {
		return clause, []any{value}
	}
	
	// Split the comma-separated string
	parts := strings.Split(strValue, ",")
	if len(parts) <= 1 {
		return clause, []any{value}
	}
	
	// Create the new placeholders
	placeholders := make([]string, len(parts))
	values := make([]any, len(parts))
	for i, part := range parts {
		placeholders[i] = "?"
		values[i] = strings.TrimSpace(part)
	}
	
	// Replace "(?)" with "(?,?,?...)"
	newClause := strings.Replace(clause, "(?)", "("+strings.Join(placeholders, ",")+")", 1)
	
	return newClause, values
}