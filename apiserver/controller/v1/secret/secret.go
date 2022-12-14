package secret

import (
	srvv1 "iam-mini/apiserver/service/v1"
	"iam-mini/apiserver/store"
)

type SecretController struct {
	srv srvv1.Service
}

func NewSecretController(store store.Factory) *SecretController {
	return &SecretController{
		srv: srvv1.NewService(store),
	}
}
