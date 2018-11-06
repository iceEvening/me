package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

//GetDB returns a DB struct
func GetDB(host string, port int, database string, username string, passwd string, maxconn int, maxidle int) *gorm.DB {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		host, port, username, passwd, database)
	for {
		db, err := gorm.Open("postgres", connInfo)
		if err != nil {
			logrus.Errorf("Get database failed: %v, retry in 10s", err)
			time.Sleep(time.Second * 10)
			continue
		}
		err = db.DB().Ping()
		if err != nil {
			logrus.Errorf("Connect to database failed: %v, retry in 10s", err)
			time.Sleep(time.Second * 10)
			continue
		}
		db.Set("gorm:table_options", "charset=utf8")
		db.DB().SetMaxOpenConns(maxconn)
		db.DB().SetMaxIdleConns(maxidle)
		return db
	}
}
