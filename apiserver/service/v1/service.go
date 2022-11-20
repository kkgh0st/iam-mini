package v1

import "iam-mini/apiserver/store"

type Service interface {
	Secrets() SecretSrv
}

type service struct {
	// service这个是个结构体，里面存储着哥哥数据结构
	store store.Factory
}

func (s *service) Secrets() SecretSrv {
	return newSecrets(s)
}

func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}
