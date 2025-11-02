# Go ORM

一个轻量级、链式调用风格的 Go 语言 ORM 库，专注于简洁易用的 SQL 构建体验。

*[English](README.md)*

## 特点与优势

- **链式 API 设计**：直观的方法链，使 SQL 构建过程清晰可读
- **零依赖**：核心功能不依赖外部库，轻量级设计
- **灵活性强**：支持模型映射、原生 SQL、子查询和 CTE 等高级功能
- **DryRun 模式**：可以预览生成的 SQL 而不实际执行，方便调试
- **全面支持 SQL 操作**：SELECT、INSERT、UPDATE、DELETE 以及各种高级查询功能
- **智能结构体映射**：Go 结构体与数据库表之间的灵活映射

## 安装

```bash
go get github.com/bobacgo/orm
```

## 快速开始

### 定义模型

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

### 基本查询示例

```go
// 查询单个模型
user := &User{}
SELECT1(user).FROM("users").WHERE(map[string]any{"AND id = ?": 1}).Query(ctx, db)

// 查询多个模型
var users []*User
SELECT2(&users).FROM("users").WHERE(map[string]any{"AND age > ?": 25}).Query(ctx, db)

// 插入数据
user := &User{Name: "张三", Age: 30}
INSERT(user).INTO("users").Exec(ctx, db)

// 更新数据
UPDATE("users").SET(map[string]any{"age": 31}).WHERE(map[string]any{"AND name = ?": "张三"}).Exec(ctx, db)

// 删除数据
DELETE().FROM("users").WHERE(map[string]any{"AND name = ?": "张三"}).Exec(ctx, db)
```

## 高级功能

### JOIN 查询

```go
SELECT("u.id", "p.product_name").
    FROM("users u").
    JOIN("products p").
    ON("u.id = p.user_id").
    Query(ctx, db)
```

### 聚合函数

```go
var count sql.Null[int64]
SELECT1(COUNT[int64]("*", &count)).FROM("users").Query(ctx, db)
```

### CTE (公共表表达式)

```go
cte := WITH("user_cte").
    AS(SELECT("id", "name").FROM("users").WHERE(map[string]any{"AND age > ?": 25}).SQL()).
    SQL()

SELECT("id", "name").FROM("user_cte").CTE(cte).Query(ctx, db)
```

### 子查询

```go
subquery := SELECT("id", "name").
    FROM("users").
    WHERE(map[string]any{"AND age > ?": 30}).
    SQL()

SELECT("id", "name").FROM("(" + subquery + ") AS old_users").Query(ctx, db)
```

## 结构体到表的映射

本 ORM 提供了 Go 结构体与数据库表之间的灵活映射系统：

```go
// 定义带有自定义字段映射的结构体
type Product struct {
    ProductID   int
    Name        string
    Description string
    Price       float64
}

// 实现 Mapping 方法以获得自定义控制
func (p *Product) Mapping() []*Mapping {
    return []*Mapping{
        {"product_id", &p.ProductID, p.ProductID},
        {"product_name", &p.Name, p.Name},
        {"description", &p.Description, p.Description},
        {"price", &p.Price, p.Price},
    }
}
```

### 映射优势

- **完全控制**：精确定义结构体字段如何映射到表列
- **双向映射**：同时支持从数据库读取和写入数据库
- **值跟踪**：同时跟踪指针和值，实现高效更新
- **自定义列名**：将结构体字段映射到不同命名的数据库列
- **灵活实现**：按照自己的方式实现 Mapping 接口

## 差异化优势

- **直观的 SQL 构建**：API 设计模仿 SQL 语法，使代码更易读易写
- **轻量级设计**：专注于 SQL 构建，不包含复杂的关系映射，更适合需要精确控制 SQL 的场景
- **灵活的查询组合**：支持子查询、CTE 等高级功能，可以构建复杂查询
- **DryRun 调试**：可以预览生成的 SQL 而不执行，方便调试和测试
- **类型安全**：利用 Go 泛型提供类型安全的查询构建和结果映射