package httpgin

//const (
//	SERVER_ERROR    = 1000 // 系统错误​
//	NOT_FOUND       = 1001 // 401错误​
//	UNKNOWN_ERROR   = 1002 // 未知错误​
//	PARAMETER_ERROR = 1003 // 参数错误​
//	AUTH_ERROR      = 1004 // 错误
//	TASKID_ERROR    = 2000
//​
//)

type APIException struct {
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
	Request   string `json:"request"`
}

func (e *APIException) Error() string {
	return e.Msg
}

func newAPIException(errorCode int, msg string) *APIException {
	return &APIException{
		ErrorCode: errorCode,
		Msg:       msg,
	}
}
