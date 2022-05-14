package orm

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitGorm(c *MySQLConfig) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database error:%+v", err))
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.LogMode(c.ShowSQL)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxIdleTime(time.Duration(c.IdleTimeout) * time.Second)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) * 2 * time.Second)
	return db
}
