package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/capomanpc/go-blog-api/api/middlewares"
)

// エラーが発生したときのレスポンス処理をここで一括で行う
// 引数はerror型であるため渡された変数がMyAppError型であったとしてもerrorインターフェースとして見られる
// 実際にMyAppError型であってもerror型として扱われるため、errors.Asを使ってMyAppError型に変換する
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	// コピー先のMyAppError型の変数を宣言
	// error型は通常ポインタで扱われる（エラーハンドリングのnilチェックなどが容易）
	var appErr *MyAppError
	// errがMyAppError型であった場合、appErrに中身をコピーする
	// このような操作によりerrorインターフェースとして扱われていたMyAppError型をMyAppError型として扱えるようになる
	if !errors.As(err, &appErr) {
		// MyAppError型に変換できない場合は想定しないエラーが発生したということなので手動でappErrに代入
		appErr = &MyAppError{
			ErrCode: Unknown, // 未知のエラー
			Message: "internal process failed",
			Err:     err,
		}
	}

	// トレースIDを取得してエラーログを出力
	traceID := middlewares.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	// appErrのErrCodeによってステータスコードを決定
	switch appErr.ErrCode {
	case NAData, BadParam:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	// ステータスコードとエラーメッセージをレスポンスとして返す
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
