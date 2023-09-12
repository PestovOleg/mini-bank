package user

import (
	"net/http"

	"github.com/PestovOleg/mini-bank/domain/user"
	"github.com/PestovOleg/mini-bank/pkg/util"
	"go.uber.org/zap"
)

type UserHandler struct {
	logger  *zap.Logger
	service *user.Service
}

func NewUserHandler(s *user.Service) *UserHandler {
	return &UserHandler{
		logger:  util.GetLogger("API"),
		service: s,
	}
}

func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
