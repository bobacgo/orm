package orm

import (
	"testing"
)

func TestDelete_SQL(t *testing.T) {
	tests := []struct {
		name  string
		build *Delete
		want  string
	}{
		{
			name:  "basic delete with where",
			build: DELETE().FROM("users").WHERE(map[string]any{"AND id = ?": 1}),
			want:  "DELETE FROM users WHERE 1 = 1 AND id = ?",
		},
		{
			name:  "delete with multiple conditions",
			build: DELETE().FROM("users").WHERE(map[string]any{"AND id = ?": 1, "AND name = ?": "Tom"}),
			want:  "DELETE FROM users WHERE 1 = 1 AND id = ? AND name = ?",
		},
		{
			name:  "table name empty",
			build: DELETE().FROM("").WHERE(map[string]any{"AND id = ?": 1}),
			want:  "DELETE FROM  WHERE 1 = 1 AND id = ?",
		},
		{
			name:  "delete with IN clause - comma-separated string",
			build: DELETE().FROM("users").WHERE(map[string]any{"id in (?)": "1,2,3"}),
			want:  "DELETE FROM users WHERE 1 = 1 id in (?,?,?)",
		},
		{
			name:  "delete with AND and IN clause",
			build: DELETE().FROM("users").WHERE(map[string]any{"AND id in (?)": "7,8,9"}),
			want:  "DELETE FROM users WHERE 1 = 1 AND id in (?,?,?)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.build.SQL()
			if got != tt.want {
				t.Errorf("SQL(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
