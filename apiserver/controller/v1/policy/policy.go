package policy

import (
	srvv1 "iam-mini/apiserver/service/v1"
	"iam-mini/apiserver/store"
)

type PolicyController struct {
	srv srvv1.Service
}

// NewPolicyController creates a policy handler.
func NewPolicyController(store store.Factory) *PolicyController {
	return &PolicyController{
		srv: srvv1.NewService(store),
	}
}
