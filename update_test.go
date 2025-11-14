package orm

import (
	"testing"
)

type testUpdateModel struct {
	ID   int
	Name string
}

func (m *testUpdateModel) Mapping() []*Mapping {
	return []*Mapping{
		{"id", &m.ID, m.ID},
		{"name", &m.Name, m.Name},
	}
}

func TestUpdate_SQL(t *testing.T) {
	tests := []struct {
		name  string
		build *Update
		want  string
	}{
		{
			name:  "basic update",
			build: UPDATE("users").SET(map[string]any{"name": "Tom"}).WHERE(map[string]any{"AND id = ?": 1}),
			want:  "UPDATE users SET name = ? WHERE id = ?",
		},
		{
			name:  "update with multiple columns",
			build: UPDATE("users").SET(map[string]any{"name": "Tom", "age": 18}).WHERE(map[string]any{"AND id = ?": 1}),
			want:  "UPDATE users SET age = ?, name = ? WHERE id = ?",
		},
		{
			name:  "update with model",
			build: UPDATE("users").SET1(&testUpdateModel{Name: "Tom"}).WHERE(map[string]any{"AND id = ?": 1}).Omit("id"),
			want:  "UPDATE users SET name = ? WHERE id = ?",
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
