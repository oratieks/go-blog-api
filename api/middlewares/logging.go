package middlewares

import (
	"log"
	"net/http"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

// newTraceID()でトレースIDを取得し、リクエスト情報をロギングするミドルウェア
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()

		// リクエスト情報をロギング
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		// トレースIDをコンテキストに格納
		// http.Request型のContextメソッドでコンテキストを取得し、SetTraceID関数でトレースIDを格納
		ctx := SetTraceID(req.Context(), traceID)

		// コンテキストを更新したリクエストを作成
		// WithContextメソッドで新しいコンテキストを持つリクエストに更新
		req = req.WithContext(ctx)

		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}

/*
	受け取ったリクエストはまずミドルウェアが受け取りミドルウェア内部で再度ServeHTTPメソッドを呼び出すことで元々のハンドラ関数を実行している
	ミドルウェア関数はfunc(http.Handler) http.Handlerとなっており、引数としてhttp.Handlerを受け取り、http.Handlerを返す関数
	ミドルウェア関数内部では、トレースIDを取得してリクエスト情報をまずロギングしている
	その後、リクエストのctxフィールドのコンテキストを取得し、内部にトレースIDを格納する
	更新したコンテキストを持つリクエストを作成し、そのリクエストを用いて元々のハンドラ関数を実行している
	こうすることでエラーが行った際にトレースIDを取得して、トレースIDと共にエラーログを出力することができる
	ここで言うエラーとはクライアントに返すためのエラーハンドラーのことである
*/
