package driver

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	DB *gorm.DB
)

type DBConfig struct {
	UserName  string
	PassWorld string
	Host      string
	Port      string
	DB        string
}

func InitMySQL(config DBConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.UserName, config.PassWorld, config.Host, config.Port, config.DB)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("[system][initMysql] open mysql error, dsn=%s", dsn)
	}
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(1000)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	DB = db
	log.Print("[system][mysql] init mysql driver success!")
}
