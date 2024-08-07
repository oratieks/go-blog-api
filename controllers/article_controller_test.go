package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestArticleListHandler(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{name: "number query", query: "1", resultCode: http.StatusOK},
		{name: "alphabet query", query: "aaa", resultCode: http.StatusNotFound},
	}

	// index, value
	for _, tt := range tests {
		// tt.nameはサブテストの名前で第二引数はテスト関数
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/article/list?page=%s", tt.query)
			// ハンドラの第二引数に渡すhttp.Request型の変数を作成（リクエストを作成）
			// 第三引数のnilはPostのときに渡したいデータを指定する
			req := httptest.NewRequest(http.MethodGet, url, nil)

			// http.ResponseRecorder型を生成
			res := httptest.NewRecorder()

			aCon.ArticleListHandler(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}
