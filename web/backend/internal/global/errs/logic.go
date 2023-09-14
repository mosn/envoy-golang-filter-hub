package errs

import (
	"fmt"
)

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"msg"`
	Origin  string `json:"origin"`
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
	return fmt.Sprintf("code:%d, msg:%s", e.Code, e.Message)
}

func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func (e *Error) WithOrigin(err error) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message,
		Origin:  fmt.Sprintf("%+v", err),
	}
}

//func (e *Error) SetLogInfo(key string, value any) *Error {
//	e.LogInfo[key] = fmt.Sprintf("%v", value)
//	return e
//}

func (e *Error) WithTips(details ...string) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message + " " + fmt.Sprintf("%v", details),
	}
	//newErr:=errors.WithMessage(e, fmt.Sprintf("%v", details))
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
