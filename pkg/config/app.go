package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

func Connect() {
	connstr := `host=localhost port=8080 user=postgres password=Pawan@2003 dbname=test_db ssl=enabled`
	d, err := gorm.Open("postgres", connstr)
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
