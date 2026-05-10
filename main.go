package main

import (
	"log"
	"net/http"
	"os" // osを再び追加

	"cloudtech-forum/handler"
	"cloudtech-forum/repository"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Println(".envファイルの読み込みに失敗しました（環境変数が直接設定されている場合は問題ありません）")
	}
}

func main() {
	// 環境変数からデータを取得
	apiport := os.Getenv("API_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// --- デバッグ用ログ: 何が読み込まれたか表示 ---
	//log.Printf("設定確認: PORT=%s, USER=%s, HOST=%s, DB=%s", port, username, host, dbname)
	// ------------------------------------------

	// データベースの接続を初期化
	err := repository.InitDB(username, password, host, port, dbname)
	if err != nil {
		log.Fatalf("データベースに接続できません: %v", err)
	}
	defer repository.CloseDB()

	// ルーティングの設定
	r := mux.NewRouter()
	r.HandleFunc("/posts", handler.CreateHandler).Methods("POST")

	// httpメソッドがGET pasthが/postsのリクエストを受け取ったときにIndexHandler関数を呼び出す
	r.HandleFunc("/posts", handler.IndexHandler).Methods("GET")

	// httpメソッドがGET pasthが/posts/idのリクエストを受け取ったときにShowHandler関数を呼び出す
	r.HandleFunc("/posts/{id:[0-9]+}", handler.ShowHandler).Methods("GET")

	// APIサーバを起動
	log.Println("APIサーバを起動しました。ポート: " + apiport)
	if err := http.ListenAndServe(":"+apiport, r); err != nil {
		log.Fatal(err)
	}

}
