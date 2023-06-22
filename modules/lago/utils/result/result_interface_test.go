package result

type ErrorCode interface {
	GetCode() string
	GetMessage() string
	GetData() interface{}
}
