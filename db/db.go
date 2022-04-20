package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var db *gorm.DB

func init() {
	initDB()
	createTableUserinfo()
}

func initDB() {
	//conn, err := gorm.Open("mysql", "root:xuxi526521258@tcp(localhost:3306)/godb")
	conn, err := gorm.Open("sqlite3", "user.db")
	if err != nil {
		panic(err)
	}
	sqlDB := conn.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 600)
	db = conn
}

type UserInfo struct {
	Id       string
	Username string
	Password string
}

func createTableUserinfo() {
	if db.HasTable(&UserInfo{}) {
		fmt.Println("exist Userinfo table")
	} else {
		fmt.Println("not exist Userinfo table")
		db.AutoMigrate(&UserInfo{})
		if !db.HasTable(&UserInfo{}) {
			fmt.Println("create Userinfo table fail!!!")
		}
	}
}

func GetUserByName(username string) UserInfo {
	user := UserInfo{}
	db.Where("username = ?", username).First(&user)
	return user
}
