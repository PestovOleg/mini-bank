package common

import "github.com/PestovOleg/mini-bank/domain/user"

type Services struct {
	UserService *user.Service
}

func NewServices(u *user.Service) *Services {
	return &Services{
		UserService: u,
	}
}
