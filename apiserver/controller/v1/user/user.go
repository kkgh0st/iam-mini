package user

import (
	srvv1 "iam-mini/apiserver/service/v1"
	"iam-mini/apiserver/store"
)

type UserController struct {
	srv srvv1.Service
}

func NewUserController(store store.Factory) *UserController {
	return &UserController{
		srv: srvv1.NewService(store),
	}
}
