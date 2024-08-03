package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

// エラーが発生したときのレスポンス処理をここで一括で行う
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError
	// 受け取ったerrをMyAppError型に入れ替える
	if !errors.As(err, &appErr) {
		// MyAppError型に変換できない場合は想定しないエラーが発生したということなので、
		// 手動でappErrに代入
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	var statusCode int

	switch appErr.ErrCode {
	case NAData, BadParam:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
