package errs

// This error code refers to the semantics of HTTP status codes
// to facilitate the identification of error types

// Type |Type description
// 1xx: Informational - Request received, continuing process
// 2xx: Success - The action was successfully received, understood, and accepted
// 3xx: Redirection - Further action must be taken in order to complete the request
// 4xx: Client Error - The request contains bad syntax or cannot be fulfilled
// 5xx: Server Error - The server failed to fulfill an apparently valid request

// 本错误码参照了HTTP状态码的语义，方便识别错误类型

// 分类 |分类描述
// 1**	信息，服务器收到请求，需要请求者继续执行操作
// 2**	成功，操作被成功接收并处理
// 3**	重定向，需要进一步的操作以完成请求
// 4**	客户端错误，请求包含语法错误或无法完成请求
// 5**	服务器错误，服务器在处理请求的过程中发生了错误

// 200 OK
var (
	success = newError(200, "Success")
)

// 400 BAD REQUEST
var (
	InvalidRequest = newError(40001, "Invalid request")
)

// 500 INTERNAL ERROR
var (
	serverInternal   = newError(50001, "Server internal error")
	Code2TokenFailed = newError(50002, "Code to token failed")
)
