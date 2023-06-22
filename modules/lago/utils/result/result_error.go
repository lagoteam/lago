package result

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type ResultError struct {
	Code    interface{} `json:"code"`
	Message string      `json:"message"`
}

func (r *ResultError) Error() string {
	err, _ := json.Marshal(r)
	return string(err)
}

func (r *ResultError) GetCode() interface{} {
	return r.Code
}

func (r *ResultError) GetMessage() string {
	return r.Message
}

func NewResultError(code interface{}, message string) error {
	return errors.Wrap(&ResultError{
		Code:    code,
		Message: message,
	}, "")
}
