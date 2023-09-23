package repository

import "github.com/jackc/pgx/v5"

// db.Find
// db.First
// db.Create
// db.Delete
// db.Save
// err := r.db.First(&product, id).Error
type db interface {
	Find(dest interface{}, conds ...interface{}) *pgx.Rows
	First(dest interface{}, conds ...interface{}) *pgx.Row

	Create(value interface{}) *pgx.Row
	Delete(value interface{}, conds ...interface{}) *pgx.Row
	Save(value interface{}) *pgx.Row
}

type Repository struct {
	con *pgx.Conn
	db  db
}
