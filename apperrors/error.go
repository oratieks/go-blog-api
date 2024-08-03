package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error `json:"-"` // エラーチェーンのため, JSONには出力しない
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
