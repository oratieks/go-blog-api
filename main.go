package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/capomanpc/go-blog-api/controllers"
	"github.com/capomanpc/go-blog-api/routers"
	"github.com/capomanpc/go-blog-api/services"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	// 最初にDBに接続することでプログラム全体でDB接続を共有する
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println(dbConn)
		log.Println("failed to connect database")
		return
	}

	// service層でdb変数を使用するのでservice層のMyAppService構造体にdb変数を渡して共有
	ser := services.NewMyAppService(db)

	// controller層
	con := controllers.NewMyAppController(ser)

	r := routers.NewRouter(con)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// .Methods(http.MethodPost)はPOSTリクエストのみを受け付けるように指定している
// "{id:[0-9]+}"は"{id}"でもok、正規表現を使いたい場合':'の後に正規表現を記述する
