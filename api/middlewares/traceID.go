package middlewares

import (
	"context"
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

// traceIDKey はコンテキストに格納するキー
// 文字列"traceID"を使うと、他のパッケージとキーが衝突する可能性がある
type traceIDKey struct{}

// 受け取ったコンテキストにWithValue関数を用いてトレースIDを格納する
func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) int {
	// コンテキストから引数のキーに対応するトレースIDを取得する
	id := ctx.Value(traceIDKey{})

	// 型アサーションid.(int)を使って、取得した値がint型かどうかを判定
	// int型のidIntを返す
	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}

// newTraceID はトレースIDを取得する
func newTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}
