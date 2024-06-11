package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // 導入 Postgres driver
)

var DB *sql.DB

func InitDB() {

	// DB 連線參數從環境變數抓取
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// fmt.Println(connStr)

	// 這邊要注意，雖然 sql.Open 看起來像是會做 DB 連線，實際上卻不會
	// 一直要等你下 SQL 的時候，才會真的連線，
	// 所以 DB 連線參數錯誤的話，這邊根本就不會有錯誤訊息 !!
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		errMsg := fmt.Sprintf("Could not connect to database: %s", err.Error())
		panic(errMsg)
	}

	// 這行是建立 connection pool 的最大數
	DB.SetMaxOpenConns(10)
	// 這行是設定沒事的時候，最多只會有5個connection保持與DB的連線
	DB.SetMaxIdleConns(5)

	// 使用 Ping 檢查連接，目的是提早發現 DB 連線問題，不要等到有 request 進來才報錯
	err = DB.Ping()
	if err != nil {
		errMsg := fmt.Sprintf("Could not connect to database (ping fail): %s", err.Error())
		panic(errMsg)
	}
}
