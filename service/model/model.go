package model

import (
	"reflect"

	"github.com/jinzhu/gorm"

	"github.com/iceEvening/me/service/model/tables"
)

//Model struct
type Model struct {
	DB *gorm.DB
}

//GetModel is a func which get a pointer of model struct
func GetModel(db *gorm.DB) *Model {
	createTables(db)
	return &Model{
		DB: db,
	}
}

func createTables(db *gorm.DB) {
	schemas := tables.Schemas{}
	for i := 0; i < reflect.TypeOf(schemas).NumField(); i++ {
		db.CreateTable(reflect.ValueOf(schemas).Field(i).Interface())
	}
}
