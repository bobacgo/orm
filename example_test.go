package orm

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"testing"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		db = new(sql.DB)
		slog.Error("open db fail", "err", err)
	}
}

func TestSelectModel(t *testing.T) {
	row := new(TestModel)

	// SELECT a, b FROM xx WHERE id = 1 GROUP BY a HAVING a > 0 ORDER BY a desc, b LIMIT 1, 1
	sqlText := SELECT(row).FROM("xx").WHERE(map[string]any{"AND id = ?": 1}).
		GROUP_BY("a").HAVING("a > 0").ORDER_BY("a desc", "b").OFFSET(1).LIMIT(1).
		SQL()
	fmt.Println(sqlText)
}

func TestSelectModels(t *testing.T) {
	var rows []*TestModel
	// SELECT a, b FROM xx WHERE id = 1 GROUP BY a HAVING a > 0 ORDER BY a desc, b LIMIT 1, 1
	sqlText := SELECT1(&rows).FROM("xx").WHERE(map[string]any{"AND id = ?": 1}).
		GROUP_BY("a").HAVING("a > 0").ORDER_BY("a desc", "b").OFFSET(1).LIMIT(1).
		SQL()
	fmt.Println(sqlText)
}

func TestSelectString(t *testing.T) {
	// SELECT a, b FROM xx WHERE id = 1 GROUP BY a HAVING a > 0 ORDER BY a desc, b LIMIT 1, 1
	sqlText := SELECT2("a", "b").FROM("xx").WHERE(map[string]any{"AND id = ?": 1}).
		GROUP_BY("a").HAVING("a > 0").ORDER_BY("a desc", "b").OFFSET(1).LIMIT(1).
		SQL()
	fmt.Println(sqlText)
}

func TestInsert(t *testing.T) {
	// INSERT INTO ab (a, b) VALUES (?, ?)
	err := INSERT1().INTO("ab").COLUMNS("a", "b").VALUES(1, 2).Print(context.Background())
	if err != nil {
		slog.Error("insert:", "err", err)
	}
}

func TestInsertModel(t *testing.T) {
	// INSERT INTO ab (a, b) VALUES (?, ?)
	err := INSERT(&TestModel{}).INTO("xxx").Print(context.Background())
	if err != nil {
		slog.Error("insert:", "err", err)
	}
}

func TestInsertModels(t *testing.T) {
	// INSERT INTO ab (a, b) VALUES (?, ?), (?, ?)
	err := INSERT(&TestModel{}, &TestModel{}).INTO("xxx").Print(context.Background())
	if err != nil {
		slog.Error("insert:", "err", err)
	}
}

func TestUpdateMap(t *testing.T) {
	// UPDATE xx SET a = 1, b = 2 WHERE a > c
	_, err := UPDATE("xx").SET(map[string]any{"a": 1, "b": 2}).WHERE(map[string]any{"a": 1}).Exec(context.Background(), db)
	if err != nil {
		slog.Error("update:", "err", err)
	}
}

func TestUpdateModel(t *testing.T) {
	// UPDATE xx SET a = 1, b = 2 WHERE a > c
	_, err := UPDATE("xx").SET1(&TestModel{A: "a", B: 1}).WHERE(map[string]any{"a": 1}).Exec(context.Background(), db)
	if err != nil {
		slog.Error("update:", "err", err)
	}
}

func TestDelete(t *testing.T) {
	// DELETE FROM xxx WHERE a = 1
	_, err := DELETE().FROM("xxx").WHERE(map[string]any{"a": 1}).Exec(context.Background(), db)
	if err != nil {
		slog.Error("delete:", "err", err)
	}
}