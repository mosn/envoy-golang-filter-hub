package errs

import "fmt"

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"msg"`
	//LogInfo map[string]string `json:"-"`
}

func newError(code int32, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
		//LogInfo: make(map[string]string),
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Message)
}

//func (e *Error) SetLogInfo(key string, value any) *Error {
//	e.LogInfo[key] = fmt.Sprintf("%v", value)
//	return e
//}

func (e *Error) WithDetails(details ...string) *Error {
	e.Message += " " + fmt.Sprintf("%v", details)
	return e
}

func (e *Error) Warp(err error) *Error {
	return e.WithDetails(err.Error())
}

//// Internal 向前端返回为内部错误，并记录原始错误
//func (e *Error) Internal(ctx context.Context, rawErr error) *Error {
//	newErr := ServerInternal.SetLogInfo("errType", e.Error())
//	logx.WithCallerSkip(1).
//		WithContext(ctx).
//		WithFields(map2LogField(newErr.LogInfo)...).
//		Error(rawErr)
//	return newErr
//}
//
//func map2LogField(m map[string]string) []logx.LogField {
//	var logFields []logx.LogField
//	for k, v := range m {
//		logFields = append(logFields, logx.LogField{
//			Key:   k,
//			Value: v,
//		})
//	}
//	return logFields
//}
