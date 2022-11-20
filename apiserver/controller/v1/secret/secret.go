package secret

import (
	srvv1 "iam-mini/apiserver/service/v1"
	"iam-mini/apiserver/store"
)

type SecretController struct {
	srv srvv1.Service
}

func NewSecretController(store store.Factory) *SecretController {
	// 这里直接返回，之后再来补充，一点也不难
	return &SecretController{
		srv: srvv1.NewService(store),
	}
}
