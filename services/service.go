package services

import "database/sql"

type MyAppService struct {
	db *sql.DB
}

func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}

// NewMyAppService構造体を返すための関数（コンストラクタ）
// dbフィールドにdb変数をセットして返す
