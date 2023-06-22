package result

type Result struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResult() *Result {
	return &Result{}
}

func Success() *Result {
	result := NewResult()
	result.SetCode(SuccessCode)                       // 200
	result.SetMessage(ResultCodeMessage[SuccessCode]) // Success
	return result
}

func Fail() *Result {
	result := NewResult()
	result.SetCode(FailCode)                       // 400
	result.SetMessage(ResultCodeMessage[FailCode]) // Fail
	return result
}

func (r *Result) SetCode(code string) *Result {
	r.Code = code
	return r
}

func (r *Result) SetMessage(message string) *Result {
	r.Message = message
	return r
}

func (r *Result) SetData(data interface{}) *Result {
	r.Data = data
	return r
}

/*func (r *Result) PutData(key string, value interface{}) *Result {
	data := map[string]interface{}{}
	if r.Data != nil {
		data = r.Data.(map[string]interface{})
	}
	data[key] = value
	r.Data = data
	return r
}*/
