package middlewares

import (
	"log"
	"net/http"
)

// レスポンス内容をログとして出力するために自分で定義したResponseWriterをハンドラ関数に渡す
type resLoggingWriter struct {
	http.ResponseWriter
	code int // レスポンスのステータスコードをここに格納してログとして出力する
}

// コンストラクタ
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// resLoggingWriter内部に含まれるhttp.ResponseWriterのWriteHeaderメソッドも使用できるが
// ステータスコードをログとして出力するためにWriteHeaderを書き換える(オーバーライドしている)
func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code                      // 引数として受け取ったcodeをcodeフィールドにコピー
	rsw.ResponseWriter.WriteHeader(code) // 元のResponseWriterのWriteHeaderメソッドを呼び出す
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// リクエスト情報をロギング
		log.Println(req.RequestURI, req.Method)

		// コンストラクタを呼び出して初期化
		rlw := NewResLoggingWriter(w)

		// 元々のハンドラに自前のresLoggingWriterを渡してハンドラを実行
		next.ServeHTTP(rlw, req)

		// コピーしておいたcodeフィールドを出力
		log.Println("res: ", rlw.code)
	})
}
