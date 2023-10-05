package db

import (
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/bloomfilter"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var (
	DB       *gorm.DB
	RDClient *redis.RDClient
	Filter   *bloomfilter.BloomFilter
)

// Init Mysql DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(conf.GetConf().GetString("db.mysql.dsn")),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	DB.AutoMigrate(&Favorite{})

	RDClient = redis.NewRDClient(conf.GetConf().GetInt("db.redis.id.favorite"))

	Filter = bloomfilter.NewBloomFilter(10000, 0.01)
}
