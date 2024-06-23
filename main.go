package main

import (
	"database/sql"
	"dbsample/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	articleId := 1000
	const sqlStr = `
		select * 
		from articles
		where article_id = ?;
	`
	row := db.QueryRow(sqlStr, articleId)
	if err := row.Err(); err != nil {
		// データ取得件数が 0 件だった場合は
		// データ読み出し処理には移らずに終了
		fmt.Println(err)
		return
	}

	var article models.Article
	var createdTime sql.NullTime
	// rows.Scan で各列のデータを article 変数に読み込みます。
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

	if err != nil {
		fmt.Println(err)
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n", article)

}
