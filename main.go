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

	const sqlStr = `
		select * from articles;
	`
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// / models.Article 型のスライスを初期化します。これはデータベースから取得した記事を格納するために使用します。
	articleArray := make([]models.Article, 0) // スライスの初期長さは0です。つまり、スライスは空（要素を持たない）状態で作成されます。
	// データベースのクエリ結果を行単位で処理します。
	for rows.Next() {
		var article models.Article
		var createdTime sql.NullTime
		// rows.Scan で各列のデータを article 変数に読み込みます。
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}

		if err != nil {
			fmt.Println(err)
		} else {
			// エラーがなければ、article 変数を articleArray スライスに追加します。
			articleArray = append(articleArray, article)
		}
	}

	fmt.Printf("%+v\n", articleArray)
}
