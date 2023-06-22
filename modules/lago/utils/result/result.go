package result

type Result struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success() *Result {
	/*return &Result{
		Code:    "200",
		Message: "Success",
	}*/
	return &Result{
		Code:    CodeSuccess,
		Message: CodeMessage[CodeSuccess],
	}
}

func Fail() *Result {
	/*return &Result{
		Code:    "400",
		Message: "Fail",
	}*/
	return &Result{
		Code:    CodeFail,
		Message: CodeMessage[CodeFail],
	}
}

func (r *Result) SetMessage(message string) *Result {
	r.Message = message
	return r
}

func (r *Result) SetCode(code interface{}) *Result {
	r.Code = code.(string)
	return r
}

func (r *Result) SetData(data interface{}) *Result {
	r.Data = data
	return r
}

func (r *Result) PutData(key string, value interface{}) *Result {
	data := map[string]interface{}{}
	if r.Data != nil {
		data = r.Data.(map[string]interface{})
	}
	data[key] = value
	r.Data = data
	return r
}
