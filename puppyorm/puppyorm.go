package puppyorm

import (
	"context"

	"gorm.io/gorm"
)

// DB is a thin wrapper over gorm.DB.
type DB struct {
	g *gorm.DB
}

// New wraps an existing *gorm.DB.
// You create the GORM connection in your app and pass it here.
func New(g *gorm.DB) *DB {
	return &DB{g: g}
}

// Query is a generic query builder for type T.
type Query[T any] struct {
	db *gorm.DB
}

// Table starts a query for a given table name with result type T.
// NOTE: This is a function, not a method, because methods
// cannot have type parameters in Go.
func Table[T any](d *DB, name string) *Query[T] {
	return &Query[T]{db: d.g.Table(name)}
}

// ==== Query builder methods ====

// Where adds a WHERE clause.
func (q *Query[T]) Where(query any, args ...any) *Query[T] {
	q.db = q.db.Where(query, args...)
	return q
}

// WhereIn: column IN (values...)
func (q *Query[T]) WhereIn(column string, values any) *Query[T] {
	q.db = q.db.Where(column+" IN ?", values)
	return q
}

// Select sets the SELECT part.
func (q *Query[T]) Select(expr any, args ...any) *Query[T] {
	q.db = q.db.Select(expr, args...)
	return q
}

// Limit adds LIMIT.
func (q *Query[T]) Limit(n int) *Query[T] {
	q.db = q.db.Limit(n)
	return q
}

// Order adds ORDER BY.
func (q *Query[T]) Order(order string) *Query[T] {
	q.db = q.db.Order(order)
	return q
}

// ==== Exec helpers ====

// Get executes the query and returns []T.
func (q *Query[T]) Get(ctx context.Context) ([]T, error) {
	var result []T
	if err := q.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// First executes the query and returns a single T.
func (q *Query[T]) First(ctx context.Context) (T, error) {
	var dest T
	err := q.db.WithContext(ctx).First(&dest).Error
	return dest, err
}

// FirstOrZero returns zero value of T if record not found.
func (q *Query[T]) FirstOrZero(ctx context.Context) (T, error) {
	var dest T
	err := q.db.WithContext(ctx).First(&dest).Error
	if err == gorm.ErrRecordNotFound {
		var zero T
		return zero, nil
	}
	return dest, err
}

// MustFirst panics on error (convenience for quick scripts / tests).
func (q *Query[T]) MustFirst(ctx context.Context) T {
	var dest T
	if err := q.db.WithContext(ctx).First(&dest).Error; err != nil {
		panic(err)
	}
	return dest
}

// Join adds an INNER JOIN using GORM's Joins().
func (q *Query[T]) Join(joinExpr string, args ...any) *Query[T] {
	// Example: "JOIN items ON items.id = orders.item_id"
	q.db = q.db.Joins(joinExpr, args...)
	return q
}

// LeftJoin adds a LEFT JOIN.
func (q *Query[T]) LeftJoin(joinExpr string, args ...any) *Query[T] {
	// GORM expects the whole JOIN clause string.
	// You can inline "LEFT JOIN ..." here for convenience.
	q.db = q.db.Joins(joinExpr, args...)
	return q
}

// RightJoin (if you really need it).
func (q *Query[T]) RightJoin(joinExpr string, args ...any) *Query[T] {
	q.db = q.db.Joins(joinExpr, args...)
	return q
}
