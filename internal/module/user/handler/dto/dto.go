package dto

//type BasicReply struct {
//	Code    int32  `json:"code" default:"200"`
//	Message string `json:"msg" default:"Success"`
//	Data    any    `json:"data"`
//}

type (
	PingRequest struct {
		Msg string `json:"msg"`
	}
	PingReply struct {
		Msg string `json:"msg"`
	}
)
