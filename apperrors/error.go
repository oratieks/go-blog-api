package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error // エラーチェーンのためのコード
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
