# puppy-go-orm 🐶

A tiny, Laravel-flavoured wrapper on top of **GORM**, combined with a **DB schema code generator**.

The goal is simple:

- ✅ Use **GORM** internally  
- ✅ Pass `*gorm.DB` from your app (no hidden DB connection logic)
- ✅ Generate **strongly-typed models** automatically from your MySQL schema  
- ✅ Write Go queries that feel almost like Laravel:

```go
order := puppyorm.Table[models.Order](db, "orders").
    Where("id = ?", 1).
    MustFirst(ctx)

fmt.Println(order.ItemId, order.CreatedAt)
🚀 Installation
Run inside your application project:

bash
Copy code
go get github.com/dilumsadeepa/puppy-go-orm@latest
You also need GORM packages in your project:

bash
Copy code
go get gorm.io/gorm
go get gorm.io/driver/mysql
🛠 Code Generator (Model Generator)
The generator reads your MySQL schema (information_schema.columns) and builds .gen.go models for every table.

Step 1 – Create a models folder
bash
Copy code
mkdir -p internal/models
Step 2 – Run the generator
bash
Copy code
go run github.com/dilumsadeepa/puppy-go-orm/cmd/puppy-go-orm-gen@latest \
  -dsn "user:password@tcp(127.0.0.1:3306)/your_db?parseTime=true" \
  -schema "your_db" \
  -out "./internal/models" \
  -pkg "models"
Flags:

Flag	Description
-dsn	Standard MySQL DSN
-schema	Your DB name
-out	Output directory for generated files
-pkg	The Go package name for generated models

Output example:
swift
Copy code
internal/models/orders.gen.go
internal/models/users.gen.go
...
Example generated struct:

go
Copy code
package models

import "time"

type Order struct {
    Id        int64     `gorm:"column:id;primaryKey" json:"id"`
    ItemId    int64     `gorm:"column:item_id" json:"item_id"`
    CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
🧩 Using puppy-go-orm in your app
Step 1 – Connect using GORM
go
Copy code
package main

import (
    "context"
    "log"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    puppyorm "github.com/dilumsadeepa/puppy-go-orm/puppyorm"
    "your-app/internal/models"
)

func main() {
    dsn := "user:password@tcp(127.0.0.1:3306)/your_db?parseTime=true"
    gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // Wrap GORM
    db := puppyorm.New(gdb)
    ctx := context.Background()

    // Query example
    order := puppyorm.Table[models.Order](db, "orders").
        Where("id = ?", 1).
        MustFirst(ctx)

    log.Println(order.ItemId, order.CreatedAt)
}
📜 Query Builder API
Table[T any](db *DB, table string) *Query[T]
Start a typed query:

go
Copy code
q := puppyorm.Table[models.Order](db, "orders")
🔎 Filters
Where(query, args...)
go
Copy code
Where("status = ?", "paid")
WhereIn(column, values)
go
Copy code
WhereIn("status", []string{"paid", "pending"})
Select(expr)
go
Copy code
Select("id, item_id, created_at")
Order(order)
go
Copy code
Order("created_at desc")
Limit(n)
go
Copy code
Limit(10)
📤 Execute
Get(ctx) ([]T, error)
Fetch multiple results:

go
Copy code
orders, err := puppyorm.
    Table[models.Order](db, "orders").
    Where("status = ?", "paid").
    Get(ctx)
First(ctx) (T, error)
Fetch one row:

go
Copy code
order, err := puppyorm.
    Table[models.Order](db, "orders").
    First(ctx)
FirstOrZero(ctx)
Returns zero-value if record not found:

go
Copy code
order, err := puppyorm.
    Table[models.Order](db, "orders").
    FirstOrZero(ctx)
MustFirst(ctx)
Panic on any error:

go
Copy code
order := puppyorm.
    Table[models.Order](db, "orders").
    MustFirst(ctx)
⚡ Optional: Auto-generate models using go generate
Inside internal/models, create file:

internal/models/doc.go

go
Copy code
package models

//go:generate go run github.com/dilumsadeepa/puppy-go-orm/cmd/puppy-go-orm-gen@latest -dsn "user:password@tcp(127.0.0.1:3306)/your_db?parseTime=true" -schema "your_db" -out "./" -pkg "models"
Then regenerate anytime:

bash
Copy code
cd internal/models
go generate ./...
📌 Roadmap
 PostgreSQL & SQLite support

 Dynamic row mode (e.g., row.String("column"))

 Custom struct/tag config

 Migrations + validation

 Query logging middleware

🐶 License
MIT License.

Use freely.
Just don’t blame the puppy if your queries are slow 🐕💤