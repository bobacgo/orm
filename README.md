# Go ORM

A lightweight, chainable Go ORM library focused on providing a clean and intuitive SQL building experience.

*[中文文档](README_zh.md)*

## Features & Advantages

- **Chainable API Design**: Intuitive method chaining for clear and readable SQL construction
- **Zero Dependencies**: Core functionality has no external dependencies
- **Highly Flexible**: Supports model mapping, raw SQL, subqueries, and CTEs
- **DryRun Mode**: Preview generated SQL without execution for easy debugging
- **Complete SQL Support**: SELECT, INSERT, UPDATE, DELETE operations with advanced query features
- **Smart Struct-to-Table Mapping**: Flexible mapping between Go structs and database tables

## Installation

```bash
go get github.com/bobacgo/orm
```

## Quick Start

### Define a Model

```go
type User struct {
    ID   int
    Name string
    Age  int
}

func (m *User) Mapping() []*Mapping {
    return []*Mapping{
        {"id", &m.ID, m.ID},
        {"name", &m.Name, m.Name},
        {"age", &m.Age, m.Age},
    }
}
```

### Basic Query Examples

```go
// Query a single model
user := &User{}
SELECT1(user).FROM("users").WHERE(map[string]any{"AND id = ?": 1}).Query(ctx, db)

// Query multiple models
var users []*User
SELECT2(&users).FROM("users").WHERE(map[string]any{"AND age > ?": 25}).Query(ctx, db)

// Insert data
user := &User{Name: "John", Age: 30}
INSERT(user).INTO("users").Exec(ctx, db)

// Update data
UPDATE("users").SET(map[string]any{"age": 31}).WHERE(map[string]any{"AND name = ?": "John"}).Exec(ctx, db)

// Delete data
DELETE().FROM("users").WHERE(map[string]any{"AND name = ?": "John"}).Exec(ctx, db)
```

## Advanced Features

### JOIN Queries

```go
SELECT("u.id", "p.product_name").
    FROM("users u").
    JOIN("products p").
    ON("u.id = p.user_id").
    Query(ctx, db)
```

### Aggregate Functions

```go
var count sql.Null[int64]
SELECT1(COUNT[int64]("*", &count)).FROM("users").Query(ctx, db)
```

### Common Table Expressions (CTE)

```go
cte := WITH("user_cte").
    AS(SELECT("id", "name").FROM("users").WHERE(map[string]any{"AND age > ?": 25}).SQL()).
    SQL()

SELECT("id", "name").FROM("user_cte").CTE(cte).Exec(ctx, db)
```

### Subqueries

```go
subquery := SELECT("id", "name").
    FROM("users").
    WHERE(map[string]any{"AND age > ?": 30}).
    SQL()

SELECT("id", "name").FROM("(" + subquery + ") AS old_users").Exec(ctx, db)
```

## Struct-to-Table Mapping

The ORM provides a flexible mapping system between Go structs and database tables:

```go
// Define a struct with custom field mapping
type Product struct {
    ProductID   int
    Name        string
    Description string
    Price       float64
}

// Implement the Mapping method for custom control
func (p *Product) Mapping() []*Mapping {
    return []*Mapping{
        {"product_id", &p.ProductID, p.ProductID},
        {"product_name", &p.Name, p.Name},
        {"description", &p.Description, p.Description},
        {"price", &p.Price, p.Price},
    }
}
```

### Mapping Advantages

- **Full Control**: Define exactly how struct fields map to table columns
- **Bidirectional Mapping**: Both reading from and writing to the database
- **Value Tracking**: Keeps track of both pointers and values for efficient updates
- **Custom Column Names**: Map struct fields to differently named database columns
- **Flexible Implementation**: Implement the Mapping interface your way

## Differentiation

- **Intuitive SQL Building**: API design mimics SQL syntax for readability
- **Lightweight Design**: Focused on SQL building without complex relationship mapping
- **Query Composition**: Support for subqueries, CTEs, and other advanced features
- **DryRun Debugging**: Preview generated SQL without execution
- **Type Safety**: Leverages Go generics for type-safe query building and result mapping
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

### TRANSACTION
```go
// Simple Transaction
// BEGIN;
// UPDATE users SET balance = 30 WHERE 1 = 1 AND name = 'Jane Doe';
// UPDATE users SET balance = 0 WHERE 1 = 1 AND name = 'John Doe';
// COMMIT;
tx := Tx(func(ctx context.Context, tx *sql.Tx) error {
    if _, err := UPDATE("users").SET(map[string]any{"balance": 30}).WHERE(map[string]any{"AND name = ?": "Jane Doe"}).Exec(ctx, tx); err != nil {return fmt.Errorf("UPDATE Jane Doe : %w", err)
        return fmt.Errorf("UPDATE Jane Doe : %w", err)
    }
    if _, err := UPDATE("users").SET(map[string]any{"balance": 0}).WHERE(map[string]any{"AND name = ?": "John Doe"}).Exec(ctx, tx); err != nil {
        return fmt.Errorf("UPDATE John Doe : %w", err)
    }
    return nil
})
if err := tx.Exec(context.Background(), db); err != nil {
    slog.Error("tx.Do", "err", err)
}
```