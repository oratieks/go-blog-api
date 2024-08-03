package services

import "errors"

// 新たなエラーを作成
var ErrNoData = errors.New("get 0 record from db.Query")
