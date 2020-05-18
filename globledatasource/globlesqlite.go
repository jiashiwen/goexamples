package globledatasource

import (
	"database/sql"
	"examples/commons"
	"github.com/gchaincl/dotsql"
	"github.com/jinzhu/gorm"
	"sync"
)

var db *gorm.DB
var l = &sync.RWMutex{}

func InitLocalStatusDB(path string) (*gorm.DB, error) {
	dbfile := path + "status.db"

	if !commons.FileExists(dbfile) {
		_, err := commons.CreateNewFile(dbfile)
		if err != nil {
			panic(err)
		}
		db, err := sql.Open("sqlite3", dbfile)
		//db, err := gorm.Open("sqlite3", dbfile)
		defer db.Close()
		if err != nil {
			panic(err)
		}
		dot, err := dotsql.LoadFromFile(path + "initdb.sql")
		if err != nil {
			panic(err)
		}
		_, err = dot.Exec(db, "create-tasks-table")
		if err != nil {
			panic(err)
		}
	}
	db, err := gorm.Open("sqlite3", dbfile)
	db.DB().SetMaxOpenConns(10)
	return db, err
}

func GetDb() *gorm.DB {
	if db == nil {
		l.Lock()
		defer l.Unlock()
		if db == nil {
			var err error
			db, err = InitLocalStatusDB("./")
			if err != nil {
				panic(err)
			}
		}
	}
	return db
}
