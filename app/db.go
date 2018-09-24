package app

import (
	"fmt"
	// MySQL driver
	"github.com/francoispqt/onelog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"github.com/zqhong/albedo/util"
	"log"
	"os"
)

var DB *Database

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

func InitDb() {
	DB = &Database{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}

// used for cli
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		Logger.ErrorWithFields(fmt.Sprintf("Database connection failed. Database name: %s", name), func(e onelog.Entry) {
			e.String("err", err.Error())
		})
		log.Printf("初始化 db 服务出错：%s\n", err.Error())
		os.Exit(1)
	}

	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	if util.IsDebug() {
		db.LogMode(true)
	}

	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用
	db.DB().SetMaxIdleConns(0)
}

func (db *Database) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}
