package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // これが必要です
)

var db *sql.DB

// データベースの初期化を行う関数
func InitDB(user, password, host, port, dbname string) (err error) { // func に修正
	// MySQLへの接続文字列を作成
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	// MySQLに接続
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Printf("データベース接続エラー: %v", err) // 引用符を修正
		return err
	}

	// データベースの接続確認を返却
	return db.Ping()
}

// データベース接続を閉じる関数
func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Printf("データベースクローズエラー: %v", err)
		}
	}
}
