package db

import "gorm.io/gorm"

type TestTable struct {
	gorm.Model
	Id   int64
	Name string
}
