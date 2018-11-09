package app

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
	"log"
	"os"
)

// 参考：http://gorm.io/docs/connecting_to_the_database.html
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
	return InitDB("db")
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return InitDB("docker_db")
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

func InitDB(dbName string) *gorm.DB {
	driver := viper.GetString(fmt.Sprintf("%s.driver", dbName))
	if driver == "" {
		driver = "mysql"
	}

	if driver == "sqlite" {
		return openSqliteDB(viper.GetString(fmt.Sprintf("%s.source", dbName)))
	}

	return openMySQLDB(viper.GetString(fmt.Sprintf("%s.username", dbName)),
		viper.GetString(fmt.Sprintf("%s.password", dbName)),
		viper.GetString(fmt.Sprintf("%s.addr", dbName)),
		viper.GetString(fmt.Sprintf("%s.name", dbName)))
}

func openSqliteDB(source string) *gorm.DB {
	db, err := gorm.Open("sqlite3", source)
	if err != nil {
		logs.Error(fmt.Sprintf("Database connection failed. Source: %s, err: %s", source, err.Error()))
		log.Printf("初始化 db 服务出错：%s\n", err.Error())
		os.Exit(1)
	}

	return db
}

func openMySQLDB(username, password, addr, name string) *gorm.DB {
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
		logs.Error(fmt.Sprintf("Database connection failed. Database name: %s, err: %s", name, err.Error()))
		log.Printf("初始化 db 服务出错：%s\n", err.Error())
		os.Exit(1)
	}

	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	if viper.GetString("gormMode") == "debug" {
		db.LogMode(true)
	}

	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用
	db.DB().SetMaxIdleConns(0)
}

func (db *Database) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}
