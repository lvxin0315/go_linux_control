package controller

const (
	SuccessMsg = "操作成功"
	ErrorMsg   = "系统异常"
)

const (
	SuccessCode = "100000"
	ErrorCode   = "100001"
)

type Output struct {
	Code   string      `json:"code" description:"编码"`
	Msg    string      `json:"msg" description:"提示信息"`
	Data   interface{} `json:"data" description:"数据"`
	Custom interface{} `json:"custom" description:"自定义信息"`
	Token  string      `json:"token" description:"Token"`
}

func (output *Output) SuccessOutput(data interface{}, msg string) *Output {
	output.Code = SuccessCode
	if msg == "" {
		output.Msg = SuccessMsg
	} else {
		output.Msg = msg
	}
	output.Data = data
	return output
}

func (output *Output) ErrorOutput(msg string) *Output {
	output.Code = ErrorCode
	if msg == "" {
		output.Msg = ErrorMsg
	}
	return output
}
