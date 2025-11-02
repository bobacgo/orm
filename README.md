# Go ORM

A simple and lightweight Object-Relational Mapping (ORM) library for Go.

## Installation

To install the ORM package, use `go get`:

```bash
go get github.com/bobacgo/orm
```

## Usage

Here are some examples of how to use the ORM:

### Model Definition

First, define a model that corresponds to your database table.

```go
type ExampleModel struct {
	ID   int
	Name string
	Age  int
}

func (m *ExampleModel) Mapping() []*Mapping {
	return []*Mapping{
		{"id", &m.ID, m.ID},
		{"name", &m.Name, m.Name},
		{"age", &m.Age, m.Age},
	}
}
```

### SELECT Queries

```go
// Simple SELECT
row := &ExampleModel{}
SELECT1(row).FROM("users").WHERE(map[string]any{"id = ?": 1}).DryRun(context.Background())

// SELECT with JOIN
SELECT("u.id", "p.product_name").FROM("users u").JOIN("products p ON u.id = p.user_id").DryRun(context.Background())

// SELECT with aggregate function
var count sql.Null[int64]
SELECT1(COUNT[int64]("*", &count)).FROM("users").DryRun(context.Background())

// SELECT with CTE (Common Table Expression)
cte := WITH("user_cte").AS(SELECT("id", "name").FROM("users").WHERE(map[string]any{"age > ?": 25}).SQL()).SQL()
SELECT("id", "name").FROM("user_cte").CTE(cte).DryRun(context.Background())

// SELECT with Subquery
SELECT("id", "name").FROM("(SELECT id, name FROM users WHERE age > 30) AS old_users").DryRun(context.Background())
```

### INSERT Queries

```go
// Simple INSERT
INSERT1().INTO("users").COLUMNS("name", "age").VALUES("John Doe", 30).DryRun(context.Background())

// INSERT from model
user := &ExampleModel{Name: "Jane Doe", Age: 25}
INSERT(user).INTO("users").DryRun(context.Background())
```

### UPDATE Queries

```go
// Simple UPDATE
UPDATE("users").SET(map[string]any{"age = ?": 31}).WHERE(map[string]any{"name = ?": "John Doe"}).DryRun(context.Background())

// UPDATE from model
user := &ExampleModel{ID: 1, Age: 32}
UPDATE("users").SET1(user).WHERE(map[string]any{"id = ?": user.ID}).DryRun(context.Background())
```

### DELETE Queries

```go
// Simple DELETE
DELETE().FROM("users").WHERE(map[string]any{"name = ?": "John Doe"}).DryRun(context.Background())
```
