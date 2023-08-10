package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/plugin/opentracing"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
)

var DB *gorm.DB

// Init Mysql DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(conf.GetConf().GetString("db.mysql.dsn")),
		&gorm.Config{
			PrepareStmt:            true,
		//	SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
}