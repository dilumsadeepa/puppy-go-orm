 puppy-go-orm Documentation  :root { --bg: #050816; --bg-card: rgba(15, 23, 42, 0.96); --bg-soft: rgba(15, 23, 42, 0.7); --accent: #22c55e; --accent-soft: rgba(34, 197, 94, 0.2); --accent-strong: #4ade80; --border-subtle: rgba(148, 163, 184, 0.25); --text-main: #e5e7eb; --text-muted: #9ca3af; --text-soft: #6b7280; --code-bg: #020617; --code-border: rgba(148, 163, 184, 0.35); --shadow-soft: 0 18px 45px rgba(15, 23, 42, 0.7); --radius-lg: 18px; --radius-pill: 999px; --transition-fast: 160ms ease-out; --font-main: system-ui, -apple-system, BlinkMacSystemFont, "SF Pro Text", "Segoe UI", sans-serif; --font-mono: ui-monospace, Menlo, Monaco, Consolas, "SF Mono", monospace; } \* { box-sizing: border-box; } body { margin: 0; padding: 0; font-family: var(--font-main); background: radial-gradient(circle at top, #1e293b 0, #020617 52%, #020617 100%); color: var(--text-main); -webkit-font-smoothing: antialiased; } .page { min-height: 100vh; display: flex; flex-direction: column; } header { position: sticky; top: 0; z-index: 20; backdrop-filter: blur(18px); background: linear-gradient( to bottom, rgba(15, 23, 42, 0.95), rgba(15, 23, 42, 0.75), transparent ); border-bottom: 1px solid rgba(148, 163, 184, 0.25); } .nav-inner { max-width: 1080px; margin: 0 auto; padding: 14px 18px 10px; display: flex; align-items: center; justify-content: space-between; gap: 16px; } .logo-wrap { display: flex; align-items: center; gap: 10px; } .logo-circle { width: 32px; height: 32px; border-radius: 999px; background: radial-gradient(circle at 30% 20%, #4ade80 0, #16a34a 35%, #22c55e 65%, #15803d 100%); display: flex; align-items: center; justify-content: center; box-shadow: 0 12px 35px rgba(34, 197, 94, 0.6); position: relative; overflow: hidden; } .logo-circle::before { content: ""; position: absolute; inset: 0; background: radial-gradient(circle at 10% -10%, rgba(255,255,255,0.4), transparent 50%); mix-blend-mode: screen; } .logo-circle span { font-size: 18px; font-weight: 800; color: #022c22; } .logo-text-main { font-weight: 650; letter-spacing: 0.03em; font-size: 18px; display: flex; align-items: baseline; gap: 4px; } .logo-text-main span:nth-child(2) { font-size: 14px; font-weight: 500; color: var(--text-soft); } .chip { border-radius: var(--radius-pill); border: 1px solid rgba(148, 163, 184, 0.25); padding: 3px 10px; display: inline-flex; align-items: center; gap: 6px; font-size: 11px; text-transform: uppercase; letter-spacing: 0.08em; color: var(--text-muted); background: radial-gradient(circle at top left, rgba(34,197,94,0.03), rgba(15,23,42,0.98)); } .chip-dot { width: 7px; height: 7px; border-radius: 999px; background: var(--accent); box-shadow: 0 0 0 4px rgba(34,197,94,0.18); } main { flex: 1 1 auto; max-width: 1080px; padding: 28px 18px 40px; margin: 0 auto; display: grid; grid-template-columns: 260px minmax(0, 1fr); gap: 24px; } @media (max-width: 860px) { main { grid-template-columns: minmax(0, 1fr); } nav.sidebar { position: static; top: auto; order: -1; } } nav.sidebar { position: sticky; top: 70px; align-self: flex-start; background: var(--bg-soft); border-radius: var(--radius-lg); border: 1px solid var(--border-subtle); padding: 12px 10px 10px; box-shadow: var(--shadow-soft); } .sidebar-title { font-size: 11px; text-transform: uppercase; letter-spacing: 0.12em; color: var(--text-soft); margin: 6px 4px 6px; } .sidebar-link { display: block; padding: 6px 10px; margin: 2px 0; border-radius: 999px; font-size: 13px; color: var(--text-muted); text-decoration: none; border: 1px solid transparent; transition: background var(--transition-fast), border-color var(--transition-fast), color var(--transition-fast), transform var(--transition-fast); } .sidebar-link:hover { background: rgba(15, 23, 42, 0.96); border-color: rgba(148, 163, 184, 0.5); color: var(--text-main); transform: translateY(-1px); } .sidebar-link strong { font-weight: 550; } .sidebar-divider { height: 1px; background: linear-gradient(to right, transparent, rgba(148,163,184,0.3), transparent); margin: 8px 0; } .content { min-width: 0; } .card { background: var(--bg-card); border-radius: var(--radius-lg); border: 1px solid var(--border-subtle); box-shadow: var(--shadow-soft); padding: 18px 16px 18px; margin-bottom: 18px; } .card + .card { margin-top: 10px; } h1, h2, h3, h4 { margin: 0 0 10px; font-weight: 650; letter-spacing: 0.02em; } h1 { font-size: 26px; display: flex; align-items: center; gap: 8px; } h1 span.sub { font-size: 13px; font-weight: 450; color: var(--text-soft); letter-spacing: 0.08em; text-transform: uppercase; } h2 { font-size: 18px; margin-top: 8px; } h3 { font-size: 16px; margin-top: 12px; } h4 { font-size: 14px; margin-top: 10px; } p { margin: 6px 0 8px; font-size: 14px; color: var(--text-muted); line-height: 1.55; } ul, ol { margin: 6px 0 8px 20px; padding: 0; font-size: 14px; color: var(--text-muted); line-height: 1.5; } li + li { margin-top: 2px; } code, pre { font-family: var(--font-mono); } code { font-size: 12px; padding: 1px 4px; border-radius: 4px; background: rgba(15,23,42,0.9); border: 1px solid rgba(148,163,184,0.4); color: var(--accent-strong); } pre { margin: 8px 0 10px; padding: 10px 12px; border-radius: 12px; background: radial-gradient(circle at top left, rgba(15,23,42,0.96), var(--code-bg)); border: 1px solid var(--code-border); overflow-x: auto; font-size: 12px; line-height: 1.55; position: relative; } pre::before { content: ""; position: absolute; inset: 0; pointer-events: none; background: radial-gradient(circle at top left, rgba(148,163,184,0.22), transparent 55%); mix-blend-mode: soft-light; opacity: 0.33; } pre code { background: transparent; border: none; color: #e5e7eb; position: relative; z-index: 1; } .badge-row { display: flex; flex-wrap: wrap; gap: 8px; margin: 4px 0 14px; } .badge { font-size: 11px; padding: 3px 8px; border-radius: 999px; border: 1px solid rgba(148,163,184,0.4); color: var(--text-soft); background: radial-gradient(circle at top left, rgba(34,197,94,0.06), rgba(15,23,42,0.96)); } .badge strong { color: var(--accent-strong); } .pill { display: inline-flex; align-items: center; gap: 6px; padding: 4px 10px; border-radius: var(--radius-pill); border: 1px solid rgba(148,163,184,0.4); background: rgba(15,23,42,0.96); font-size: 11px; color: var(--text-muted); text-transform: uppercase; letter-spacing: 0.08em; } .pill span { font-size: 11px; } .pill small { opacity: 0.8; } .section-header { display: flex; align-items: baseline; justify-content: space-between; gap: 10px; } .section-header .hint { font-size: 11px; color: var(--text-soft); } .callout { border-radius: 12px; padding: 8px 10px; border: 1px solid rgba(148,163,184,0.35); background: radial-gradient(circle at top left, rgba(34,197,94,0.06), rgba(15,23,42,0.96)); font-size: 12px; color: var(--text-muted); display: flex; gap: 8px; align-items: flex-start; margin: 6px 0 10px; } .callout strong { color: var(--accent-strong); } .callout-icon { font-size: 13px; margin-top: 1px; } .table-like { border-radius: 12px; border: 1px solid rgba(148,163,184,0.35); overflow: hidden; margin: 8px 0 10px; } .table-like-row { display: grid; grid-template-columns: 130px minmax(0, 1fr); font-size: 13px; } .table-like-row:nth-child(odd) { background: rgba(15,23,42,0.95); } .table-like-row:nth-child(even) { background: rgba(2,6,23,0.98); } .table-like-row div { padding: 6px 10px; } .table-like-row div:first-child { border-right: 1px solid rgba(148,163,184,0.24); color: var(--text-soft); font-weight: 500; } .table-like-row div:last-child { color: var(--text-muted); } footer { padding: 12px 18px 18px; font-size: 12px; color: var(--text-soft); text-align: center; } a { color: var(--accent-strong); text-decoration: none; } a:hover { text-decoration: underline; } .inline-list { margin: 0; padding: 0; list-style: none; display: flex; flex-wrap: wrap; gap: 8px; font-size: 12px; } .inline-list li { color: var(--text-soft); }

🐶

puppy-go-orm for GORM

**Go** + Laravel-ish query builder Codegen models from MySQL schema

Docs v0.x

Contents

[**1\. Overview** · What & why](#intro) [**2\. Install** · Add to project](#install) [**3\. Model Generator** · From DB](#models) [**4\. Basic Usage** · Table queries](#usage) [**5\. Joins & Raw** · Complex SQL](#joins) [**6\. API Reference** · Query methods](#api) [**7\. go generate** · Auto models](#generate)

[**Examples** · Snippets](#examples)

puppy-go-orm Laravel-ish ORM layer on top of GORM
=================================================

**puppy-go-orm** is a tiny wrapper on top of `gorm.io/gorm` that gives you:

*   Typed, generic query builder: `Table[T](db, "orders")`
*   A **model generator** that reads your MySQL schema and writes Go structs for each table
*   Readable, Laravel-flavoured chainable methods over GORM
*   Typed results even for **raw SQL with joins**

    // Example: get a single order by ID
    
    order := puppyorm.Table[models.Order](db, "orders").
        Where("id = ?", 1).
        MustFirst(ctx)
    
    fmt.Println(order.ItemId, order.CreatedAt)

💡

Think of `Table[T]` as `DB::table('orders')` in Laravel, but strongly-typed and compiled.

2\. Installation
----------------

Add to your Go project

### 2.1 Add the module

From your project root:

    go get github.com/dilumsadeepa/puppy-go-orm@latest
    go mod tidy

### 2.2 Install GORM in your app (if not already)

    go get gorm.io/gorm
    go get gorm.io/driver/mysql    # or postgres/sqlite etc.

### 2.3 Wrap your existing GORM DB

    import (
        "gorm.io/driver/mysql"
        "gorm.io/gorm"
    
        puppyorm "github.com/dilumsadeepa/puppy-go-orm/puppyorm"
    )
    
    func newDB() (*puppyorm.DB, error) {
        dsn := "user:pass@tcp(127.0.0.1:3306)/your_db?parseTime=true"
    
        gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err != nil {
            return nil, err
        }
    
        return puppyorm.New(gdb), nil
    }

3\. Model Generator
-------------------

Generate Go structs from MySQL

The generator reads `information_schema.columns` and creates a Go file per table with GORM + JSON tags.

### 3.1 Create models folder

    mkdir -p internal/models

### 3.2 Run the generator (all tables)

    go run github.com/dilumsadeepa/puppy-go-orm/cmd/puppy-go-orm-gen@latest \
      -dsn "user:pass@tcp(127.0.0.1:3306)/your_db?parseTime=true" \
      -schema "your_db" \
      -out "./internal/models" \
      -pkg "models"

### 3.3 Generate only specific tables

    go run github.com/dilumsadeepa/puppy-go-orm/cmd/puppy-go-orm-gen@latest \
      -dsn "user:pass@tcp(127.0.0.1:3306)/your_db?parseTime=true" \
      -schema "your_db" \
      -out "./internal/models" \
      -pkg "models" \
      -tables "orders,users,reservations"

This will create files like:

    internal/models/orders.gen.go
    internal/models/users.gen.go
    ...

Example generated struct:

    package models
    
    import "time"
    
    type Order struct {
        Id        int64     `gorm:"column:id;primaryKey" json:"id"`
        ItemId    int64     `gorm:"column:item_id" json:"item_id"`
        CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
    }

4\. Basic Usage
---------------

Query tables like Laravel, with types

### 4.1 Get a single row

    import (
        "context"
    
        puppyorm "github.com/dilumsadeepa/puppy-go-orm/puppyorm"
        "your-app/internal/models"
    )
    
    func example(db *puppyorm.DB) error {
        ctx := context.Background()
    
        order := puppyorm.Table[models.Order](db, "orders").
            Where("id = ?", 1).
            MustFirst(ctx)
    
        fmt.Println(order.ItemId, order.CreatedAt)
        return nil
    }

### 4.2 Get a list (like `get()` in Laravel)

    orders, err := puppyorm.Table[models.Order](db, "orders").
        Where("status = ?", "paid").
        Order("created_at DESC").
        Limit(20).
        Get(ctx)
    
    if err != nil {
        // handle
    }
    
    for _, o := range orders {
        fmt.Println(o.Id, o.ItemId)
    }

### 4.3 Pluck a single column

    var ids []int64
    
    err := puppyorm.Table[models.Order](db, "orders").
        Where("status = ?", "paid").
        Pluck(ctx, "id", &ids)
    
    if err != nil {
        // handle
    }
    
    fmt.Println(ids) // []int64{...}

5\. Joins & Raw SQL
-------------------

For complex reporting queries

### 5.1 Joins with a custom DTO struct

Define a struct shaped like your SELECT:

    type OrderWithItem struct {
        OrderID   int64   `gorm:"column:order_id"`
        ItemID    int64   `gorm:"column:item_id"`
        ItemTitle string  `gorm:"column:item_title"`
        CityName  *string `gorm:"column:city_name"`
    }

Then use `Select` + `Join`:

    rows, err := puppyorm.
        Table[OrderWithItem](db, "orders o").
        Select(`
            o.id        AS order_id,
            o.item_id   AS item_id,
            i.title     AS item_title,
            c.name      AS city_name
        `).
        Join("JOIN reservations r ON r.order_id = o.id").
        Join("JOIN delivery_area da ON da.id = r.delivery_area_id").
        Join("JOIN cities c ON c.id = da.city_id").
        Where("o.status = ?", "paid").
        Get(ctx)

### 5.2 Raw SQL with typed results

Sometimes, using full SQL is easier. You still get types:

    type OrderReportRow struct {
        OrderID   int64   `gorm:"column:order_id"`
        CityName  string  `gorm:"column:city_name"`
        Total     float64 `gorm:"column:total"`
    }
    
    rows, err := puppyorm.
        Raw[OrderReportRow](db, `
            SELECT
                o.id        AS order_id,
                c.name      AS city_name,
                o.total     AS total
            FROM orders o
            JOIN reservations r ON r.order_id = o.id
            JOIN cities c       ON c.id = r.city_id
            WHERE o.created_at BETWEEN ? AND ?
            ORDER BY o.id DESC
        `, from, to).
        Get(ctx)

6\. API Reference
-----------------

Main types & methods

### 6.1 Core types

`type DB`

Wrapper around `*gorm.DB`. Created via `New(*gorm.DB)`.

`type Query[T any]`

Generic query builder for result type `T`.

### 6.2 Constructors

*   `New(g *gorm.DB) *DB`
*   `Table[T any](db *DB, table string) *Query[T]`
*   `Raw[T any](db *DB, sql string, args ...any) *Query[T]`

### 6.3 Filters (Where)

*   `Where(query any, args ...any) *Query[T]`
*   `OrWhere(query any, args ...any) *Query[T]`
*   `WhereIn(column string, values any) *Query[T]`
*   `WhereNotIn(column string, values any) *Query[T]`
*   `WhereNull(column string) *Query[T]`
*   `WhereNotNull(column string) *Query[T]`
*   `WhereBetween(column string, from, to any) *Query[T]`
*   `WhereNotBetween(column string, from, to any) *Query[T]`

### 6.4 Joins

*   `Join(joinExpr string, args ...any) *Query[T]`
*   `LeftJoin(joinExpr string, args ...any) *Query[T]`
*   `RightJoin(joinExpr string, args ...any) *Query[T]`

### 6.5 Select & ordering

*   `Select(expr any, args ...any) *Query[T]`
*   `Distinct() *Query[T]`
*   `GroupBy(columns ...string) *Query[T]`
*   `Having(query any, args ...any) *Query[T]`
*   `Order(order string) *Query[T]` & `OrderDesc(col string)`
*   `Limit(n int) *Query[T]` & `Offset(n int) *Query[T]`

### 6.6 Execution & aggregates

*   `Get(ctx context.Context) ([]T, error)` — like Laravel `get()->toArray()`
*   `First(ctx context.Context) (T, error)`
*   `FirstOrZero(ctx context.Context) (T, error)`
*   `MustFirst(ctx context.Context) T` — panics on error
*   `Count(ctx context.Context) (int64, error)`
*   `Exists(ctx context.Context) (bool, error)`
*   `Pluck(ctx context.Context, column string, dest any) error`

### 6.7 Write helpers

*   `Create(ctx context.Context, value *T) error`
*   `UpdateColumns(ctx context.Context, values map[string]any) error`
*   `Delete(ctx context.Context) error`

7\. Using `go generate`
-----------------------

Regenerate models automatically

Add a small file in your models package:

    // internal/models/doc.go
    package models
    
    //go:generate go run github.com/dilumsadeepa/puppy-go-orm/cmd/puppy-go-orm-gen@latest -dsn "user:pass@tcp(127.0.0.1:3306)/your_db?parseTime=true" -schema "your_db" -out "./" -pkg "models"

Then from `internal/models`:

    cd internal/models
    go generate ./...

Any time your DB schema changes, run this to refresh structs.

Examples
--------

Common snippets

### Example: Report query with raw SQL

    type OrderSummary struct {
        OrderID   int64   `gorm:"column:order_id"`
        CityName  string  `gorm:"column:city_name"`
        Total     float64 `gorm:"column:total"`
    }
    
    func GetOrderSummary(ctx context.Context, db *puppyorm.DB, from, to string) ([]OrderSummary, error) {
        return puppyorm.Raw[OrderSummary](db, `
            SELECT
                o.id    AS order_id,
                c.name  AS city_name,
                o.total AS total
            FROM orders o
            JOIN reservations r ON r.order_id = o.id
            JOIN cities c       ON c.id = r.city_id
            WHERE o.created_at BETWEEN ? AND ?
            ORDER BY o.id DESC
        `, from, to).Get(ctx)
    }

### Example: Check if any "flagged" orders exist

    exists, err := puppyorm.
        Table[models.Order](db, "orders").
        Where("flag = ?", 1).
        Exists(ctx)
    
    if err != nil {
        // handle
    }
    if exists {
        fmt.Println("We have flagged orders")
    }

*   puppy-go-orm 🐶
*   ·
*   Wrapper around `gorm.io/gorm`
*   ·
*   Inspired by Laravel's query builder