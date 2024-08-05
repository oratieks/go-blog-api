package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NAData           ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam            ErrCode = "R002"
)

// ErrCode型のWrapメソッドを定義

// レシーバはポインタレシーバではなく、値レシーバになっている
// ポインタレシーバではメソッドが呼び出される際に、インスタンスのポインタをメソッドに渡す
// そのためインスタンスを直接変更することができる
// 値レシーバではメソッドが呼び出される際に、インスタンスのコピーをメソッドに渡す
// そのためインスタンスを直接変更することができない

// 定数のアドレスへは言語レベルでアクセスできないため、ポインタレシーバにするとコンパイルエラーになる
// （メソッドが呼び出された際に、定数のアドレスを渡すことができないため）
// （メソッドを呼び出す際には、そのメソッドに対してレシーバで指定されている型のインスタンスと引数を渡す必要がある）
func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
