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

	/*
		使用 docker compose 來啟動環境時，DB 還沒來的及 ready，go web api 這邊就先 try 導致直接退出。
		理論上 docker compose 那邊可以額外寫 script 來判斷 DB 到底好了沒，不過還沒時間研究，
		目前暫時解決方式是，將 go web api 這邊的 compose 設定成 restart: always
	*/

	// 使用 Ping 檢查連接，目的是提早發現 DB 連線問題，不要等到有 request 進來才報錯
	err = DB.Ping()
	if err != nil {
		errMsg := fmt.Sprintf("Could not connect to database (ping fail): %s", err.Error())
		panic(errMsg)
	}

	// 如果沒建立 Table，就先直接幫它建
	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
        id      BIGSERIAL,
        name    TEXT,
        PRIMARY KEY (id)
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create users table.")
	}
}
