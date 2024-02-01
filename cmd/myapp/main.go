package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres" // Docker Composeで定義したPostgreSQLサービスの名前に変更
	port     = 5432
	user     = "yourusername" // 実際のPostgreSQLのユーザー名に変更
	password = "yourpassword" // 実際のPostgreSQLのパスワードに変更
	dbname   = "yourdbname"   // 実際のデータベース名に変更
)

func main() {
	// データベースへの接続
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// データベース接続の確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// テストデータの書き込み
	name := "John Doe"
	age := 30
	_, err = db.Exec("INSERT INTO test_table (name, age) VALUES ($1, $2)", name, age)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully!")
}
