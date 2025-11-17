package puppyorm

import (
	"context"

	"gorm.io/gorm"
)

// DB is a thin wrapper over gorm.DB.
// You create *gorm.DB in your app and pass it into New().
type DB struct {
	g *gorm.DB
}

// New wraps an existing *gorm.DB.
func New(g *gorm.DB) *DB {
	return &DB{g: g}
}

// Query is a Laravel-flavoured query builder over GORM
// for a specific result type T.
type Query[T any] struct {
	db *gorm.DB
}

// Table starts a query for the given table name.
// Example: Table[models.Order](db, "orders")
func Table[T any](d *DB, table string) *Query[T] {
	return &Query[T]{db: d.g.Table(table)}
}

// Raw starts a query from a raw SQL string, still returning a typed Query[T].
// Example: Raw[MyDTO](db, "SELECT ... FROM ... WHERE x = ?", 123)
func Raw[T any](d *DB, sql string, args ...any) *Query[T] {
	return &Query[T]{db: d.g.Raw(sql, args...)}
}

// Underlying returns the underlying *gorm.DB (escape hatch).
func (q *Query[T]) Underlying() *gorm.DB {
	return q.db
}

//
// ---------- “Where” / filter methods (Laravel-like) ----------
//

// Where adds a basic WHERE clause.
// Equivalent to Laravel: ->where('status', 'paid') or ->where('id', 1)
func (q *Query[T]) Where(query any, args ...any) *Query[T] {
	q.db = q.db.Where(query, args...)
	return q
}

// OrWhere adds an OR WHERE.
// Laravel: ->orWhere('status', 'pending')
func (q *Query[T]) OrWhere(query any, args ...any) *Query[T] {
	q.db = q.db.Or(query, args...)
	return q
}

// WhereIn adds "column IN (?)".
// Laravel: ->whereIn('id', [1,2,3])
func (q *Query[T]) WhereIn(column string, values any) *Query[T] {
	q.db = q.db.Where(column+" IN ?", values)
	return q
}

// WhereNotIn adds "column NOT IN (?)".
// Laravel: ->whereNotIn('id', [1,2,3])
func (q *Query[T]) WhereNotIn(column string, values any) *Query[T] {
	q.db = q.db.Where(column+" NOT IN ?", values)
	return q
}

// WhereNull adds "column IS NULL".
func (q *Query[T]) WhereNull(column string) *Query[T] {
	q.db = q.db.Where(column + " IS NULL")
	return q
}

// WhereNotNull adds "column IS NOT NULL".
func (q *Query[T]) WhereNotNull(column string) *Query[T] {
	q.db = q.db.Where(column + " IS NOT NULL")
	return q
}

// WhereBetween adds "column BETWEEN ? AND ?".
func (q *Query[T]) WhereBetween(column string, from, to any) *Query[T] {
	q.db = q.db.Where(column+" BETWEEN ? AND ?", from, to)
	return q
}

// WhereNotBetween adds "column NOT BETWEEN ? AND ?".
func (q *Query[T]) WhereNotBetween(column string, from, to any) *Query[T] {
	q.db = q.db.Where(column+" NOT BETWEEN ? AND ?", from, to)
	return q
}

//
// ---------- Joins (inner/left/right) ----------
//

// Join adds a generic JOIN clause.
// Example: Join("JOIN items i ON i.id = orders.item_id")
func (q *Query[T]) Join(joinExpr string, args ...any) *Query[T] {
	q.db = q.db.Joins(joinExpr, args...)
	return q
}

// LeftJoin adds a LEFT JOIN clause.
// Example: LeftJoin("LEFT JOIN items i ON i.id = orders.item_id")
func (q *Query[T]) LeftJoin(joinExpr string, args ...any) *Query[T] {
	q.db = q.db.Joins(joinExpr, args...)
	return q
}

// RightJoin adds a RIGHT JOIN clause.
func (q *Query[T]) RightJoin(joinExpr string, args ...any) *Query[T] {
	q.db = q.db.Joins(joinExpr, args...)
	return q
}

//
// ---------- Select / group / ordering / limits ----------
//

// Select sets the SELECT clause.
// Laravel: ->select('id', 'name')
func (q *Query[T]) Select(expr any, args ...any) *Query[T] {
	q.db = q.db.Select(expr, args...)
	return q
}

// Distinct adds DISTINCT.
func (q *Query[T]) Distinct() *Query[T] {
	q.db = q.db.Distinct()
	return q
}

// GroupBy sets GROUP BY.
// Laravel: ->groupBy('status')
func (q *Query[T]) GroupBy(columns ...string) *Query[T] {
	for _, c := range columns {
		q.db = q.db.Group(c)
	}
	return q
}

// Having adds HAVING clause.
func (q *Query[T]) Having(query any, args ...any) *Query[T] {
	q.db = q.db.Having(query, args...)
	return q
}

// Order sets ORDER BY.
// Laravel: ->orderBy('created_at', 'desc')
func (q *Query[T]) Order(order string) *Query[T] {
	q.db = q.db.Order(order)
	return q
}

// OrderDesc is a helper: ORDER BY col DESC.
func (q *Query[T]) OrderDesc(column string) *Query[T] {
	q.db = q.db.Order(column + " DESC")
	return q
}

// Limit sets LIMIT.
func (q *Query[T]) Limit(n int) *Query[T] {
	q.db = q.db.Limit(n)
	return q
}

// Offset sets OFFSET.
func (q *Query[T]) Offset(n int) *Query[T] {
	q.db = q.db.Offset(n)
	return q
}

//
// ---------- Exec / result methods ----------
//

// Get executes the query and returns a slice of T.
//
// This is effectively equivalent to Laravel's: ->get()->toArray()
// because in Go, []T *is* the array.
func (q *Query[T]) Get(ctx context.Context) ([]T, error) {
	var result []T
	if err := q.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// First returns the first row that matches.
//
// Laravel: ->first()
func (q *Query[T]) First(ctx context.Context) (T, error) {
	var dest T
	err := q.db.WithContext(ctx).First(&dest).Error
	return dest, err
}

// FirstOrZero returns zero value of T if no record is found.
// In Laravel terms, like first() but instead of null you get a zero-value struct.
func (q *Query[T]) FirstOrZero(ctx context.Context) (T, error) {
	var dest T
	err := q.db.WithContext(ctx).First(&dest).Error
	if err == gorm.ErrRecordNotFound {
		var zero T
		return zero, nil
	}
	return dest, err
}

// MustFirst is a convenience: panic on any error.
// Handy in scripts / tests where you want short code.
func (q *Query[T]) MustFirst(ctx context.Context) T {
	var dest T
	if err := q.db.WithContext(ctx).First(&dest).Error; err != nil {
		panic(err)
	}
	return dest
}

// Count returns COUNT(*) for the query.
//
// Laravel: ->count()
func (q *Query[T]) Count(ctx context.Context) (int64, error) {
	var c int64
	if err := q.db.WithContext(ctx).Count(&c).Error; err != nil {
		return 0, err
	}
	return c, nil
}

// Exists checks whether any row exists for the query.
//
// Laravel: ->exists()
func (q *Query[T]) Exists(ctx context.Context) (bool, error) {
	// Efficient pattern: SELECT 1 LIMIT 1
	var dest T
	err := q.db.WithContext(ctx).Limit(1).Find(&dest).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

//
// ---------- Pluck / aggregation helpers ----------
//

// Pluck fills dest with a single column's values.
//
// Usage (like Laravel's ->pluck('id')):
//
//	var ids []int64
//	err := puppyorm.Table[models.Order](db, "orders").
//	    Where("status = ?", "paid").
//	    Pluck(ctx, "id", &ids)
//
// dest must be a pointer to a slice, e.g. *[]int64, *[]string, etc.
func (q *Query[T]) Pluck(ctx context.Context, column string, dest any) error {
	return q.db.WithContext(ctx).Pluck(column, dest).Error
}

//
// ---------- Write helpers (optional, Eloquent-ish) ----------
//

// Create inserts the given value.
//
// Laravel: Model::create([...])
func (q *Query[T]) Create(ctx context.Context, value *T) error {
	return q.db.WithContext(ctx).Create(value).Error
}

// UpdateColumns performs an update on matched rows with the given map.
//
// Laravel: ->update([...])
func (q *Query[T]) UpdateColumns(ctx context.Context, values map[string]any) error {
	return q.db.WithContext(ctx).Updates(values).Error
}

// Delete deletes matched rows.
//
// Laravel: ->delete()
func (q *Query[T]) Delete(ctx context.Context) error {
	// GORM uses table set by Table(...), model type is new(T).
	return q.db.WithContext(ctx).Delete(new(T)).Error
}
