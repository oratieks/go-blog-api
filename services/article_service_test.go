package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/capomanpc/go-blog-api/services"
	_ "github.com/go-sql-driver/mysql"
)

var aSer *services.MyAppService

func TestMain(m *testing.M) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aSer = services.NewMyAppService(db)

	m.Run()
}

func BenchmarkGetArticleService(b *testing.B) {
	// 取得する記事番号を決定
	articleID := 1

	// 計測している時間をリセット（前処理の時間が入らないように）
	b.ResetTimer()

	// 繰り返し回数はb.Nで決められる
	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			// エラーが発生したらテストが失敗するようにb.Errorを呼び出す
			b.Error(err)
			break
		}
	}
}
