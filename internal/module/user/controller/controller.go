package controller

type IUserController interface {
	Ping(req PingRequest) (PingReply, error)
}

type UserController struct {
}

func NewUserController() IUserController {
	return UserController{}
}
