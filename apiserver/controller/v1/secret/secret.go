package secret

import (
	"iam-mini/apiserver/store"
)

type SecretController struct {
	// 这里有service，但是我们这里省略掉
}

func NewSecretController(store store.Factory) *SecretController {
	// 这里直接返回，之后再来补充，一点也不难
	return &SecretController{}
}
