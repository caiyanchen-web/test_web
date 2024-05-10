package msg

// 用于返回对应函数的错误信息以及错误吗
type Msg struct {
	Code int
	Msg  string
	Data interface{}
}

func NewMsg(code int, msg string, date interface{}) *Msg {
	return &Msg{
		Code: code,
		Msg:  msg,
		Data: date,
	}
}
