package orm

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func InitXorm(c *MySQLConfig) (db *xorm.Engine) {
	db, err := xorm.NewEngine("mysql", c.DSN)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database error:%+v", err))
	}

	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	db.ShowSQL(c.ShowSQL)
	db.SetConnMaxLifetime(time.Duration(c.IdleTimeout))
	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("failed to ping database error:%+v", err))
	}
	return db
}

func SyncVBS9010(db *xorm.Engine) error {
	return db.Sync2(
		&MxAccount{},
		&MxAdminUser{},
		&MxApp{},
		&MxAppExtra{},
		&MxAppPanel{},
		&MxAppVersion{},
		&MxBinding{},
		&MxCity{},
		&MxCompany{},
		&MxCountry{},
		&MxDevice{},
		&MxDeviceOnline{},
		&MxDeviceRule{},
		&MxDeviceStat{},
		&MxDoc{},
		&MxEno{},
		&MxFaq{},
		&MxFeedback{},
		&MxFeedbackMsg{},
		&MxHome{},
		&MxHomeRule{},
		&MxLoginLog{},
		&MxMenu{},
		&MxProduct{},
		&MxProductCategory{},
		&MxProject{},
		&MxProvince{},
		&MxPush{},
		&MxPushRecord{},
		&MxRole{},
		&MxTeam{},
		&MxToken{},
		&MxUser{},
		&MxUserClient{},
		&MxUserInfo{},
		&MxUserStat{},
		&MxUserThird{},
		&MxVirtual{},
	)
}
